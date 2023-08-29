package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDownLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M80 104v96l-48-48Z" opacity=".2"/><path d="M224 48a8 8 0 0 0-8 8a88.1 88.1 0 0 1-88 88H88v-40a8 8 0 0 0-13.66-5.66l-48 48a8 8 0 0 0 0 11.32l48 48A8 8 0 0 0 88 200v-40h40A104.11 104.11 0 0 0 232 56a8 8 0 0 0-8-8ZM72 180.69L43.31 152L72 123.31Z"/></g>`),
		g.Group(children),
	)
}