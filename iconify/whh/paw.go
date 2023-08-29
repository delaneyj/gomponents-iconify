package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1002 398.5q-26 49.5-66 71t-70 2.5t-32.5-67.5t23.5-98t65.5-71T992 233t33 67.5t-23 98zM513 865q-58 0-86.5-.5t-73.5-4t-67.5-11.5t-47.5-22t-35-36.5t-10-53.5q0-42 16-94t46-102.5t68.5-93T412 379t101-26t101 26t88.5 68.5t68.5 93T817 643t16 94q0 31-10 53.5T788 827t-47.5 22t-67.5 11.5t-73.5 4t-86.5.5zm124-578.5Q600 276 584.5 227t-.5-108t51.5-93T709 2.5T761.5 62t1 108t-51.5 93t-74 23.5zm-248 0Q352 297 315.5 263T264 170t.5-108T317 2.5T390.5 26t51.5 93t-.5 108t-52.5 59.5zm-365 112q-26-49.5-23-98T34 233t69.5 2.5t65.5 71t23.5 98T160 472t-70-2.5t-66-71z"/>`),
		g.Group(children),
	)
}