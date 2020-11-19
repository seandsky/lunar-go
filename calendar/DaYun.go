package calendar

import "github.com/6tail/lunar-go/LunarUtil"

type DaYun struct {
	// 开始年(含)
	startYear int
	// 结束年(含)
	endYear int
	// 开始年龄(含)
	startAge int
	// 结束年龄(含)
	endAge int
	// 序数，0-9
	index int
	// 运
	yun   *Yun
	lunar *Lunar
}

func NewDaYun(yun *Yun, index int) *DaYun {
	daYun := new(DaYun)
	daYun.yun = yun
	daYun.lunar = yun.GetLunar()
	daYun.index = index
	birthYear := yun.GetLunar().GetSolar().GetYear()
	year := yun.GetStartSolar().GetYear()
	if daYun.index < 1 {
		daYun.startYear = birthYear
		daYun.startAge = 1
		daYun.endYear = year - 1
		daYun.endAge = year - birthYear
	} else {
		add := (index - 1) * 10
		daYun.startYear = year + add
		daYun.startAge = daYun.startYear - birthYear + 1
		daYun.endYear = daYun.startYear + 9
		daYun.endAge = daYun.startAge + 9
	}
	return daYun
}

func (daYun *DaYun) GetStartYear() int {
	return daYun.startYear
}

func (daYun *DaYun) GetEndYear() int {
	return daYun.endYear
}

func (daYun *DaYun) GetStartAge() int {
	return daYun.startAge
}

func (daYun *DaYun) GetEndAge() int {
	return daYun.endAge
}

func (daYun *DaYun) GetIndex() int {
	return daYun.index
}

func (daYun *DaYun) GetLunar() *Lunar {
	return daYun.lunar
}

// 获取干支
func (daYun *DaYun) GetGanZhi() string {
	if daYun.index < 1 {
		return ""
	}
	offset := LunarUtil.GetJiaZiIndex(daYun.lunar.GetMonthInGanZhiExact())
	if daYun.yun.IsForward() {
		offset += daYun.index
	} else {
		offset -= daYun.index
	}

	size := len(LunarUtil.JIA_ZI)
	if offset >= size {
		offset -= size
	}
	if offset < 0 {
		offset += size
	}
	return LunarUtil.JIA_ZI[offset]
}

// 获取所在旬
func (daYun *DaYun) GetXun() string {
	return LunarUtil.GetXun(daYun.GetGanZhi())
}

// 获取旬空(空亡)
func (daYun *DaYun) GetXunKong() string {
	return LunarUtil.GetXunKong(daYun.GetGanZhi())
}

// 获取流年
func (daYun *DaYun) GetLiuNian() []*LiuNian {
	n := 10
	if daYun.index < 1 {
		n = daYun.endYear - daYun.startYear + 1
	}
	l := make([]*LiuNian, n)
	for i := 0; i < n; i++ {
		l[i] = NewLiuNian(daYun, i)
	}
	return l
}

// 获取小运
func (daYun *DaYun) GetXiaoYun() []*XiaoYun {
	n := 10
	if daYun.index < 1 {
		n = daYun.endYear - daYun.startYear + 1
	}
	l := make([]*XiaoYun, n)
	for i := 0; i < n; i++ {
		l[i] = NewXiaoYun(daYun, i, daYun.yun.IsForward())
	}
	return l
}
