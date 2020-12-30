package common

// ArrayOf does the array contain specified item
func ArrayOf(arr []string, dest string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == dest {
			return true
		}
	}
	return false
}

// ArrayDuplice 数组去重
func ArrayDuplice(arr []string) []string {
	var out []string
	tmp := make(map[string]byte)
	for _, v := range arr {
		tmplen := len(tmp)
		tmp[v] = 0
		if len(tmp) != tmplen {
			out = append(out, v)
		}
	}
	return out
}