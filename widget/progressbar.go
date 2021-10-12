package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 进度条
type ProgressBar struct {
	Element
}

// 进度条_创建, 创建进度条元素
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口
func NewProgressBar(x int, y int, cx int, cy int, hParent int) *ProgressBar {
	p := &ProgressBar{}
	p.SetHandle(xc.XProgBar_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象
func NewProgressBarByHandle(handle int) *ProgressBar {
	p := &ProgressBar{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil
func NewProgressBarByName(name string) *ProgressBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil
func NewProgressBarByUID(nUID int) *ProgressBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil
func NewProgressBarByUIDName(name string) *ProgressBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 进度条_置范围, 设置范围
// range_: 范围.
func (p *ProgressBar) SetRange(range_ int) int {
	return xc.XProgBar_SetRange(p.Handle, range_)
}

// 进度条_取范围
func (p *ProgressBar) GetRange() int {
	return xc.XProgBar_GetRange(p.Handle)
}

// 进度条_置进度图片
// hImage: 图片句柄.
func (p *ProgressBar) SetImageLoad(hImage int) int {
	return xc.XProgBar_SetImageLoad(p.Handle, hImage)
}

// 进度条_置两端大小, 设置两端间隔大小
// leftSize: 左边间隔大小.
// rightSize: 右边间隔大小.
func (p *ProgressBar) SetSpaceTwo(leftSize int, rightSize int) int {
	return xc.XProgBar_SetSpaceTwo(p.Handle, leftSize, rightSize)
}

// 进度条_置进度, 设置位置点
// pos: 位置点.
func (p *ProgressBar) SetPos(pos int) int {
	return xc.XProgBar_SetPos(p.Handle, pos)
}

// 进度条_取进度, 获取当前位置点
func (p *ProgressBar) GetPos() int {
	return xc.XProgBar_GetPos(p.Handle)
}

// 进度条_置水平, 设置水平或垂直
// bHorizon: 水平或垂直.
func (p *ProgressBar) SetHorizon(bHorizon bool) int {
	return xc.XProgBar_SetHorizon(p.Handle, bHorizon)
}

/*
以下都是事件
*/

type XE_PROGRESSBAR_CHANGE func(pos int, pbHandled *bool) int            // 进度条元素,进度改变事件.
type XE_PROGRESSBAR_CHANGE1 func(hEle int, pos int, pbHandled *bool) int // 进度条元素,进度改变事件.

// 进度条元素,进度改变事件.
func (p *ProgressBar) Event_PROGRESSBAR_CHANGE(pFun XE_PROGRESSBAR_CHANGE) bool {
	return xc.XEle_RegEventC(p.Handle, xcc.XE_PROGRESSBAR_CHANGE, pFun)
}

// 进度条元素,进度改变事件.
func (p *ProgressBar) Event_PROGRESSBAR_CHANGE1(pFun XE_PROGRESSBAR_CHANGE1) bool {
	return xc.XEle_RegEventC1(p.Handle, xcc.XE_PROGRESSBAR_CHANGE, pFun)
}
