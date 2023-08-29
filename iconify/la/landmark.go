package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3.906L3.625 9.062L3 9.345V12h2v11H3v5h26v-5h-2V12h2V9.344l-.625-.281zm0 2.188L25.375 10H6.625zM7 12h2v11H7zm4 0h2v11h-2zm4 0h2v11h-2zm4 0h2v11h-2zm4 0h2v11h-2zM5 25h22v1H5z"/>`),
		g.Group(children),
	)
}