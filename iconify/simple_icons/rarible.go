package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rarible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.8 0A4.79 4.79 0 0 0 0 4.8v14.4A4.79 4.79 0 0 0 4.8 24h14.4a4.79 4.79 0 0 0 4.8-4.8V4.8A4.79 4.79 0 0 0 19.2 0zm1.32 7.68h8.202c2.06 0 3.666.44 3.666 2.334c0 1.137-.671 1.702-1.427 1.898c.904.268 1.558 1 1.558 2.16v2.131h-3.451V14.18c0-.62-.37-.87-1-.87H9.572v2.893H6.12zm3.452 2.5v.834h4.155c.452 0 .726-.06.726-.416c0-.358-.274-.418-.726-.418z"/>`),
		g.Group(children),
	)
}