package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlitchLab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.884 34.73a8.197 8.197 0 0 0 9.76-9.76l6.562-6.562a16.774 16.774 0 0 1-22.884 22.885Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.782 35.833a16.772 16.772 0 1 1 22.884-22.885l-6.562 6.562a8.191 8.191 0 1 0-9.76 9.76Z"/>`),
		g.Group(children),
	)
}