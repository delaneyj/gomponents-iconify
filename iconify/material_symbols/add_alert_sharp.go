package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAlertSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-2h2v-7q0-2.075 1.25-3.688T10.5 4.2V2h3v2.2q2 .5 3.25 2.113T18 10v7h2v2H4Zm8 3q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm-1-7h2v-2h2v-2h-2V9h-2v2H9v2h2v2Z"/>`),
		g.Group(children),
	)
}