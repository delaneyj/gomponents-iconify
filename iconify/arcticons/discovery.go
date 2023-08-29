package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discovery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="16.239" cy="28.435" r="11.739" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.24 16.696v-8.87h11.086A16.174 16.174 0 0 1 43.5 24h0a16.174 16.174 0 0 1-16.174 16.174H16.24"/>`),
		g.Group(children),
	)
}