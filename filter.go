package content

// CollectionFilterResults contains the results of a filter operation
type CollectionFilterResults interface {
	Purpose() string
	Original() Collection
	Filtered() Collection
	Errors() []error
}

// CollectionFilterItemFn returns the content and true / error if a specific item should be filtered
type CollectionFilterItemFn func(index int) (Content, bool, error)

// CollectionFilterRangeFn returns the range and filter item function
type CollectionFilterRangeFn func() (int, int, CollectionFilterItemFn)

// filterResults implements both FilteredCollection and Collection interface contracts
type filterResults struct {
	purpose  string
	original Collection
	filtered filteredCollection
	errors   []error
}

type filteredCollection []Content

func (f filterResults) Purpose() string {
	return f.purpose
}

func (f filterResults) Original() Collection {
	return f.original
}

func (f filterResults) Filtered() Collection {
	return f
}

func (f filterResults) Source() string {
	return f.original.Source()
}

func (f filterResults) Content() ([]Content, error) {
	return f.filtered, nil
}

func (f filterResults) Errors() []error {
	return f.errors
}

// MakeFilteredCollection returns a
func MakeFilteredCollection(purpose string, original Collection, rangeFn CollectionFilterRangeFn) CollectionFilterResults {
	result := new(filterResults)
	result.purpose = purpose
	result.original = original
	start, end, keepFn := rangeFn()
	for i := start; i <= end; i++ {
		c, ok, err := keepFn(i)
		if ok {
			result.filtered = append(result.filtered, c)
		}
		if err != nil {
			result.errors = append(result.errors, err)
		}
	}
	return result
}
