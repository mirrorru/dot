package dot

// SliceToSlice - converts slice to new type slice
func SliceToSlice[FROM, TO any](source []FROM, converter func(FROM) TO) []TO {
	if source == nil {
		return nil
	}

	result := make([]TO, len(source))
	for i := range source {
		result[i] = converter(source[i])
	}

	return result
}

func SliceToSliceError[FROM, TO any](source []FROM, converter func(FROM) (TO, error)) (result []TO, err error) {
	if source == nil {
		return nil, nil
	}

	result = make([]TO, len(source))
	for i := range source {
		result[i], err = converter(source[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
