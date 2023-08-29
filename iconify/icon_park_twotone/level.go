package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Level(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLevel0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 42L4 18.5L9.695 6h28.61L44 18.5L24 42Z"/><path d="m32 18l-8 9l-8-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLevel0)"/>`),
		g.Group(children),
	)
}