package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningOther(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 20h12v2H18zm0 4h12v2H18zm0 4h12v2H18zm-4-10a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 14 18zM13 7h2v9h-2z"/><path fill="currentColor" d="M14 4a10.011 10.011 0 0 1 10 10h2a12 12 0 1 0-12 12v-2a10 10 0 0 1 0-20Z"/>`),
		g.Group(children),
	)
}