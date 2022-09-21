package layouts

import (
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type fieldTags struct {
	Column              string
	CommaSeparatedValue bool
	Email               bool
	Required            bool
	Regex               string
	DateFormat          string
	Max                 float64
	Min                 float64
	MaxLength           int64
	MinLength           int64
	Url                 bool
	Unique              bool
	hasMin              bool
	hasMax              bool
	hasMinLength        bool
	hasMaxLength        bool
}

func parseOptions(tags string) (fieldTags, error) {
	ft := fieldTags{}
	tags = strings.TrimSpace(tags)

	if len(tags) == 0 {
		return ft, ErrTagNoFieldTag
	}

	name := "excelLayout:"
	res := strings.Index(tags, name)
	if res < 0 {
		return ft, ErrTagNoFieldTag
	}
	tags = tags[res+len(name):]
	if len(tags) == 0 {
		return ft, ErrTagEmptyFieldTag
	}

	res = strings.Index(tags, "\"")
	if res < 0 {
		return ft, ErrTagEmptyFieldTag
	}
	tags = tags[res+1:]

	res = strings.Index(tags, "\"")
	if res < 0 {
		return ft, ErrTagEmptyFieldTag
	}
	tags = strings.TrimSpace(tags[:res])
	if len(tags) < 1 {
		return ft, ErrTagEmptyFieldTag
	}
	options := strings.Split(tags, ",")

	if len(options) == 0 {
		return ft, ErrTagEmptyFieldTag
	}

	for _, o := range options {
		pair := strings.SplitN(o, ":", 2)
		key := strings.ToLower(strings.TrimSpace(pair[0]))
		val := ""
		if len(pair) > 1 {
			val = strings.TrimSpace(pair[1])
		}

		switch key {
		case "column":
			if val == "" {
				return ft, ErrTagMissingColumnValue
			}
			ft.Column = strings.TrimSpace(strings.ToUpper(pair[1]))
		case "commaseparatedvalue":
			ft.CommaSeparatedValue = true
		case "dateformat":
			if val == "" {
				return ft, ErrTagMissingDateFormatValue
			}
			ft.DateFormat = strings.TrimSpace(pair[1])
		case "regex":
			if val == "" {
				return ft, ErrTagMissingRegexValue
			}
			ft.Regex = strings.TrimSpace(pair[1])
		case "email":
			ft.Email = true
		case "required":
			ft.Required = true
		case "max":
			if val == "" {
				return ft, ErrTagMissingMaxValue
			}
			ft.hasMax = true
			v, _ := strconv.ParseFloat(val, 32)
			ft.Max = v
		case "min":
			if val == "" {
				return ft, ErrTagMissingMinValue
			}
			ft.hasMin = true
			v, _ := strconv.ParseFloat(val, 32)
			ft.Min = v
		case "maxlength":
			if val == "" {
				return ft, ErrTagMissingMaxValue
			}
			ft.hasMaxLength = true
			v, _ := strconv.ParseInt(val, 0, 32)
			ft.MaxLength = v
		case "minlength":
			if val == "" {
				return ft, ErrTagMissingMaxValue
			}
			ft.hasMinLength = true
			v, _ := strconv.ParseInt(val, 0, 32)
			ft.MinLength = v
		case "url":
			ft.Url = true
		case "unique":
			ft.Unique = true
		}

	}

	if ft.Regex != "" {
		if _, err := regexp.Compile(ft.Regex); err != nil {
			return ft, ErrRegexInvalid
		}
	}

	if (ft.hasMax && ft.hasMin) && (ft.Max < ft.Min) {
		return ft, ErrTagInvalidMaxMinValues
	}

	if (ft.hasMaxLength && ft.hasMinLength) && (ft.MaxLength < ft.MinLength) {
		return ft, ErrTagInvalidMaxMinLengthValues
	}

	return ft, nil
}

func parseStringRules(v string, tags fieldTags) (string, []error) {
	value := strings.TrimSpace(v)
	errors := []error{}
	if tags.Required && len(value) == 0 {
		errors = append(errors, ErrRequiredValueRuleFail)
	}
	if tags.hasMin {
		errors = append(errors, ErrTagMinForbidden)
	}
	if tags.hasMax {
		errors = append(errors, ErrTagMaxForbidden)
	}
	if len(value) > 0 {

		if tags.hasMinLength && (int(tags.MinLength) > len(value)) {
			errors = append(errors, ErrMinLengthValueRuleFail)
		}
		if tags.hasMaxLength && (int(tags.MaxLength) < len(value)) {
			errors = append(errors, ErrMaxLengthValueRuleFail)
		}
		if tags.Url {
			if _, err := url.ParseRequestURI(value); err != nil {
				errors = append(errors, ErrUrlValueRuleFail)
			}
		}
		if tags.Email {
			if _, err := mail.ParseAddress(value); err != nil {
				errors = append(errors, ErrEmailValueRuleFail)
			}
		}
		if tags.Regex != "" {
			regex, err := regexp.Compile(tags.Regex)
			if err != nil {
				errors = append(errors, ErrRegexInvalid)
			}
			if match := regex.MatchString(value); !match {
				errors = append(errors, ErrRegexRuleFail)
			}
		}
	}
	if len(errors) > 0 {
		return "", errors
	}
	return value, nil
}

func parseIntRules(v string, tags fieldTags) (int64, []error) {
	value := strings.TrimSpace(v)
	errors := []error{}
	if tags.Required && value == "" {
		errors = append(errors, ErrRequiredValueRuleFail)
	} else if !tags.Required && value == "" {
		value = "0"
	}
	if tags.hasMinLength {
		errors = append(errors, ErrTagMinLengthForbidden)
	}
	if tags.hasMaxLength {
		errors = append(errors, ErrTagMaxLengthForbidden)
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		errors = append(errors, ErrIntegerInvalid)
		return 0, errors
	}
	if tags.hasMin && int(tags.Min) > val {
		errors = append(errors, ErrMinValueRuleFail)
	}
	if tags.hasMax && int(tags.Max) < val {
		errors = append(errors, ErrMaxValueRuleFail)
	}
	if len(errors) > 0 {
		return 0, errors
	}
	return int64(val), nil
}

func parseFloat64Rules(v string, tags fieldTags) (float64, []error) {
	value := strings.TrimSpace(v)
	errors := []error{}
	if tags.Required && value == "" {
		errors = append(errors, ErrRequiredValueRuleFail)
	} else if !tags.Required && value == "" {
		value = "0"
	}
	if tags.hasMinLength {
		errors = append(errors, ErrTagMinLengthForbidden)
	}
	if tags.hasMaxLength {
		errors = append(errors, ErrTagMaxLengthForbidden)
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		errors = append(errors, ErrDecimalInvalid)
		return 0, errors
	}
	if tags.hasMin && tags.Min > val {
		errors = append(errors, ErrMinValueRuleFail)
	}
	if tags.hasMax && tags.Max < val {
		errors = append(errors, ErrMaxValueRuleFail)
	}

	if len(errors) > 0 {
		return 0.0, errors
	}

	return val, nil
}

func parseTimeRules(v string, tags fieldTags) (time.Time, []error) {
	errs := []error{}
	var date time.Time

	if tags.Required && v == "" {
		errs = append(errs, ErrRequiredValueRuleFail)
	}
	if !tags.Required && v == "" {
		return date, errs
	}
	if tags.DateFormat == "" {
		tags.DateFormat = "2006-01-02 03:04:05"
	}

	regex := regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`)
	match := regex.MatchString(v)
	if match {
		in, _ := strconv.ParseFloat(v, 64)
		excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
		date = excelEpoch.Add(time.Duration(in * float64(24*time.Hour)))
	} else if parsed, err := time.Parse(tags.DateFormat, v); err == nil {
		date = parsed
	} else {
		errs = append(errs, ErrDateFormatInvalid)
	}

	if len(errs) == 0 {
		return date, nil
	}

	return date, errs
}
