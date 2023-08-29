package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="30.5" cy="24" r="6.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.5 30.5A6.5 6.5 0 1 0 24 37v-6.5h-6.5m13-13a6.5 6.5 0 1 0 0-13H24v13h6.5Zm-13-13a6.5 6.5 0 1 0 0 13H24v-13h-6.5Zm0 13a6.5 6.5 0 1 0 0 13H24v-13h-6.5Z"/>`),
		g.Group(children),
	)
}