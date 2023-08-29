package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0C4.478 0 0 4.478 0 10s4.478 10 10 10s10-4.478 10-10S15.522 0 10 0Zm0 18.304A8.305 8.305 0 0 1 3.56 4.759l11.681 11.68A8.266 8.266 0 0 1 10 18.305Zm6.44-3.063L4.759 3.561a8.305 8.305 0 0 1 11.68 11.68Z"/>`),
		g.Group(children),
	)
}