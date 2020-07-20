package slices

func Interface2String(iArr []interface{}) (sArr []string, err error) {
	sArr = make([]string, len(iArr))
	var ok bool
	for i, v := range iArr {
		sArr[i], ok = v.(string)
		if !ok {
			return nil, err
		}
	}
	return sArr, nil
}

func String2Interface(sArr []string) (iArr []interface{}) {
	iArr = make([]interface{}, len(sArr))
	for i, v := range sArr {
		iArr[i] = v
	}
	return iArr
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}