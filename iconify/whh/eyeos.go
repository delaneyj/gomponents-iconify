package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyeos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 512q0-53 37.5-90.5T512 384t90.5 37.5T640 512t-37.5 90.5T512 640t-90.5-37.5T384 512zm544 128q-40 0-68-28t-28-68q0-50-13.5-108T783 342q-43-69-115-109.5T512 192q-87 0-160.5 43T235 351.5T192 512t43 160.5T351.5 789T512 832q40 0 78.5-10t62.5-22t47.5-22t35.5-10q40 0 68 28t28 68q0 28-10 49.5T802 942l-9 8q-1 0-1 .5t-1 .5q-125 73-279 73q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0q173 0 309.5 104T1004 371q20 69 20 173q0 40-28 68t-68 28z"/>`),
		g.Group(children),
	)
}