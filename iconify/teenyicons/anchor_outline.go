package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 4.497a2 2 0 0 0 2-1.998a2 2 0 0 0-4 0a2 2 0 0 0 2 1.998Zm0 0v9.994m0 0c-3.866 0-7-3.132-7-6.996h2m5 6.996c3.866 0 7-3.132 7-6.996h-2"/>`),
		g.Group(children),
	)
}