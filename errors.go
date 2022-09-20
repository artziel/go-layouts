package layouts

import "errors"

// ExcelLayout Errors
var ErrExcelNoSheetFound error = errors.New("no sheet found on file")
var ErrExcelValidationFail error = errors.New("file rows validation fail")

// Parser Errors
var ErrTagNoFieldTag error = errors.New("no \"excelLayout\" tag found")
var ErrTagEmptyFieldTag error = errors.New("empty \"excelLayout\" tag found")
var ErrTagMissingColumnValue error = errors.New("expected value for \"column\" tag entry")
var ErrTagMissingRegexValue error = errors.New("expected value for \"regex\" tag entry")
var ErrTagMissingDateFormatValue error = errors.New("expected value for \"dateformat\" tag entry")
var ErrTagMissingMaxValue error = errors.New("expected value for \"max\" tag entry")
var ErrTagMissingMinValue error = errors.New("expected value for \"min\" tag entry")
var ErrTagInvalidMaxMinValues error = errors.New("the \"max\" value should be greater than \"min\" value tag entry")
var ErrTagInvalidMaxMinLengthValues error = errors.New("the \"maxLength\" value should be greater than \"minLength\" value tag entry")
var ErrTagMissingMinLengthValue error = errors.New("expected value for \"minLength\" tag entry")
var ErrTagMissingMaxLengthValue error = errors.New("expected value for \"maxLength\" tag entry")
var ErrTagMinForbidden error = errors.New("the use of value \"min\" tag entry is not allow for strings")
var ErrTagMaxForbidden error = errors.New("the use of value \"max\" tag entry is not allow for strings")
var ErrTagMinLengthForbidden error = errors.New("the use of value \"minLength\" tag entry is not allow for numbers")
var ErrTagMaxLengthForbidden error = errors.New("the use of value \"maxLength\" tag entry is not allow for numbers")
var ErrRequiredValueRuleFail error = errors.New("value required rule fail")
var ErrMinValueRuleFail error = errors.New("min value rule fail")
var ErrMaxValueRuleFail error = errors.New("max value rule fail")
var ErrMinLengthValueRuleFail error = errors.New("min length rule fail")
var ErrMaxLengthValueRuleFail error = errors.New("max length rule fail")
var ErrUrlValueRuleFail error = errors.New("url value rule validation fail")
var ErrEmailValueRuleFail error = errors.New("email value rule validation fail")
var ErrRegexRuleFail error = errors.New("regex matching rule fail")
var ErrRegexInvalid error = errors.New("invalid regex value")
var ErrDateFormatInvalid error = errors.New("invalid or unexpected datetime format")
var ErrIntegerInvalid error = errors.New("invalid integer value")
var ErrDecimalInvalid error = errors.New("invalid integer value")
var ErrCommaSeparatedInvalid error = errors.New("invalid comma separated expected value")
var ErrNotUnique error = errors.New("value is not unique")
