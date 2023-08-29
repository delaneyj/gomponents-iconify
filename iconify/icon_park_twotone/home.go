package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHome0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M9 18v24h30V18L24 6L9 18Z"/><path fill="#555" stroke-linejoin="round" d="M19 29v13h10V29H19Z"/><path stroke-linecap="round" d="M9 42h30"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHome0)"/>`),
		g.Group(children),
	)
}