package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M16 15a3 3 0 0 1 3-3h18a3 3 0 0 1 3 3v18a3 3 0 0 1-3 3H19a3 3 0 0 1-3-3V15Z"/><path stroke-linecap="round" d="M8 4v40m0-25h8M8 29h8m6-5h12m-6-6v12"/></g>`),
		g.Group(children),
	)
}