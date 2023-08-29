package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNotebook0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 6a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v36a2 2 0 0 1-2 2H12a2 2 0 0 1-2-2V6Z"/><path stroke="#000" stroke-linecap="round" d="M34 6v36"/><path stroke="#fff" stroke-linecap="round" d="M6 14h8M6 24h8M6 34h8M27 4h12M27 44h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNotebook0)"/>`),
		g.Group(children),
	)
}