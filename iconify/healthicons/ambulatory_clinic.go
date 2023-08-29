package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmbulatoryClinic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 6a6.002 6.002 0 0 1 5.653 3.985H42.5V12.5L38 15.532V40h3a1 1 0 1 1 0 2h-3v.015H10V42H7a1 1 0 1 1 0-2h3V15.497l-4-3.03V9.984h7.347A6.002 6.002 0 0 1 19 6Zm1 5V9h-2v2h-2v2h2v2h2v-2h2v-2h-2Zm14 13v16h-8V24h8Zm-13 0h-7v7h7v-7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}