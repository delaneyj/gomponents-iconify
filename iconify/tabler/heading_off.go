package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadingOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12h5m4 0h1M7 7v12M17 5v8m0 4v2m-2 0h4M15 5h4M5 19h4M3 3l18 18"/>`),
		g.Group(children),
	)
}