package helper

func GetPaginations(count, limit, page int) (offset, totalPage int) {
	totalPage = ((count - 1) / limit) + 1
	offset = limit * (page - 1)
	return
}
