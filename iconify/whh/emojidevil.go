package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emojidevil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M969.032 282q55 110 55 231q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5t-199-40.5t-163.5-109t-109-163.5t-40.5-199q0-121 55-231l-48-218q-14-38 2-54.5t54-2.5l147 93q135-99 301.5-99t302.5 99l147-93q37-14 53.5 2.5t2.5 54.5zm-329 103q0 53 28 90.5t68 37.5t68-37.5t28-90.5q0-47-24-84zm-128 512q67 0 139.5-19t126.5-57.5t54-83.5q0-13-9.5-22.5t-22.5-9.5h-576q-13 0-22.5 9.5t-9.5 22.5q0 45 54 83.5t126.5 57.5t139.5 19zm-224-384q40 0 68-37.5t28-90.5l-168-84q-24 37-24 84q0 53 28 90.5t68 37.5z"/>`),
		g.Group(children),
	)
}