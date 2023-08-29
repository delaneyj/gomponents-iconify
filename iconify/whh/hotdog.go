package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotdog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1008.342 743l-727-727q12-12 40.5-14.5t61.5 4.5t68.5 25t60.5 43l438 438q25 25 43 62t25.5 70.5t5 61.5t-15.5 37zm-20 245q-36 36-87 36t-87-36l-778-778q-36-36-36-87t36-87t87-36t87 36l778 778q36 36 36 87t-36 87zm-245 19q-9 14-37 16.5t-61.5-5t-70.5-25.5t-62-43l-437-438q-25-25-43.5-62t-25.5-71t-4.5-61.5t15.5-37.5z"/>`),
		g.Group(children),
	)
}