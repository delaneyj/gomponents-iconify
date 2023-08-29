package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUDownLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M80 120v96l-48-48Z" opacity=".2"/><path d="M168 48H80a8 8 0 0 0 0 16h88a48 48 0 0 1 0 96H88v-40a8 8 0 0 0-13.66-5.66l-48 48a8 8 0 0 0 0 11.32l48 48A8 8 0 0 0 88 216v-40h80a64 64 0 0 0 0-128ZM72 196.69L43.31 168L72 139.31Z"/></g>`),
		g.Group(children),
	)
}