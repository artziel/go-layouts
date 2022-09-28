package layouts

import "sync"

type Error struct {
	RowIndex int
	Error    error
	Column   string
	Value    string
}

type Row struct {
	Index int
}

type Layout struct {
	Rows        interface{}
	uniques     map[string]map[string]int
	errors      []Error
	IgnoreEmpty bool
	errLock     *sync.Mutex
}

func (l *Layout) GetErrors() []Error {
	l.errLock.Lock()
	defer l.errLock.Unlock()
	return l.errors
}

func (l *Layout) HasErrors() bool {
	l.errLock.Lock()
	defer l.errLock.Unlock()
	return len(l.errors) > 0
}

func (l *Layout) AddError(err Error) {
	l.errLock.Lock()
	defer l.errLock.Unlock()
	l.errors = append(l.errors, err)
}

func (l *Layout) AddErrors(errs []Error) {
	l.errLock.Lock()
	defer l.errLock.Unlock()
	l.errors = append(l.errors, errs...)
}

func (l *Layout) AddRow(r interface{}) error {
	if l.Rows == nil {
		l.Rows = []interface{}{}
	}

	rows := l.Rows.([]interface{})
	l.Rows = append(rows, r)

	return nil
}

func (l *Layout) GetRows() []interface{} {
	return l.Rows.([]interface{})
}

func (l *Layout) Iterate(fnc func(i int, r interface{}) error) error {

	for i, r := range l.GetRows() {
		fnc(i, r)
	}

	return nil
}

func (l *Layout) CountRows() int {
	rows := l.Rows.([]interface{})
	return len(rows)
}

func NewLayout(t interface{}) Layout {

	l := Layout{
		Rows:    []interface{}{},
		errors:  []Error{},
		uniques: map[string]map[string]int{},
		errLock: &sync.Mutex{},
	}
	return l
}
