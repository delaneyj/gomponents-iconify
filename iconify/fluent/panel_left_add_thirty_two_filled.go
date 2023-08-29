package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftAddThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 28h9.015a8.968 8.968 0 0 1-1.003-2H12V6h13.5A2.5 2.5 0 0 1 28 8.5v7.015a9.05 9.05 0 0 1 2 1.828V8.5A4.5 4.5 0 0 0 25.5 4h-19A4.5 4.5 0 0 0 2 8.5v15A4.5 4.5 0 0 0 6.5 28ZM23 30.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm1-12.25V22h3.75a.75.75 0 0 1 0 1.5H24v3.75a.75.75 0 0 1-1.5 0V23.5h-3.75a.75.75 0 0 1 0-1.5h3.75v-3.75a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}