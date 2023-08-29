package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModalList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 6h2v1H3V6zm3 0h7v1H6V6zM3 8h2v1H3V8zm3 0h7v1H6V8zm-3 2h2v1H3v-1zm3 0h7v1H6v-1z"/><path fill="currentColor" d="M0 1v14h16V1H0zm15 13H1V4h14v10zm0-11h-1V2h1v1z"/>`),
		g.Group(children),
	)
}