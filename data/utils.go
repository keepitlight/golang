package data

// Chunk continuous call f until it returns false or error, size is the page size >= 1,
// start and end are the data record position offset, incremented.
// Note: ⚠️Use for statement whenever possible, and pay attention to boundary conditions, and to avoid dead loops
//
// 持续调用 f 直到其返回 false 或 error，size 为分页大小，start 和 end 为 0 开始的数据记录位置偏移，自增。
// 注意，⚠️尽量使用 for 循环语句，注意边界条件，避免死循环
func Chunk(size int, f func(start, end int) (next bool, err error)) (count int, err error) {
	if size < 1 {
		// 不执行
		return
	}
	for {
		var next bool
		if next, err = f(count*size, (count+1)*size); err != nil {
			return
		} else if !next {
			break
		}
		count++
	}
	return
}

// Offset continuous call f until its return false or error, parameter limit is the limit of offset capacity.
// Note: ⚠️Use for statement whenever possible, and pay attention to boundary conditions, and to avoid dead loops
//
// 偏移，持续调用 f 直到其返回 false 或 error，参数 limit 为限制的偏移容量，offset 为 0 起的位置偏移值，自增。
// 注意，⚠️尽量使用 for 循环语句，注意边界条件，避免死循环
func Offset(limit int, f func(offset int) (next bool, err error)) (count int, err error) {
	if limit < 1 {
		// 不执行
		return
	}
	for {
		var next bool
		if next, err = f(count * limit); err != nil {
			return
		} else if !next {
			break
		}
		count++
	}
	return
}

// Paged continuous call f until its return false or error, parameter capacity is the page size of data records,
// number is the 1 start page number, incremented.
// Note: ⚠️Use for statement whenever possible, and pay attention to boundary conditions, and to avoid dead loops
//
// 分页，持续调用 f 直到其返回 false 或 error，参数 size 为每页数据记录容量，number 为 1 起的页号，自增，
// start 为 0 开始的数据记录偏移位置，自增，end 为 start + size
// 注意，⚠️尽量使用 for 循环语句，注意边界条件，避免死循环
func Paged(capacity int, f func(number, start, end int) (next bool, err error)) (count int, err error) {
	if capacity < 1 {
		// 不执行
		return 0, nil
	}
	for {
		count++
		var next bool
		if next, err = f(count, (count-1)*capacity, count*capacity); err != nil {
			return
		} else if !next {
			break
		}
	}
	return
}
