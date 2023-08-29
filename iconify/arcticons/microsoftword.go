package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoftword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M5.5 16v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2h-16a2 2 0 0 0-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9 18.5l3.25 11m3.25-11l-3.25 11m3.25-11l3.25 11m3.25-11l-3.25 11M13.23 14V6.5a2 2 0 0 1 2-2H40.5a2 2 0 0 1 2 2v35a2 2 0 0 1-2 2H15.23a2 2 0 0 1-2-2V34"/>`),
		g.Group(children),
	)
}