package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtonPass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.242" height="31.242" x="8.379" y="8.379" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4" transform="rotate(45 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.624 39.716L4.736 26.828a4.002 4.002 0 0 1 0-5.657L17.624 8.283a3.001 3.001 0 0 1 4.243 0l13.595 13.595a3.001 3.001 0 0 1 0 4.243L21.867 39.716a3.001 3.001 0 0 1-4.243 0Z"/>`),
		g.Group(children),
	)
}