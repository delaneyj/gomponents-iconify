package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileJpg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileJpg0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M10 4h20l10 10v28a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M31 18H21m5 0v12"/><path stroke="#000" stroke-linecap="round" d="M18 30a4 4 0 0 0 8 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileJpg0)"/>`),
		g.Group(children),
	)
}