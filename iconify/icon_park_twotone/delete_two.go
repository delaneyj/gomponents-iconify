package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDeleteTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M14 11L4 24l10 13h30V11H14Z"/><path d="m21 19l10 10m0-10L21 29"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDeleteTwo0)"/>`),
		g.Group(children),
	)
}