package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.16 2a9.844 9.844 0 0 1-4.149 1a9.968 9.968 0 0 1-4.229-1.026C1.171 2 .001 3.17.001 5.84v6A2.19 2.19 0 0 0 2.191 14h.232a2.19 2.19 0 0 0 2.074-1.485C4.802 11.6 5.642 10 6.582 10h2.84c.94 0 1.78 1.6 2.08 2.5A2.194 2.194 0 0 0 13.58 14h.232c1.21 0 2.19-.98 2.19-2.19v-6c0-2.64-1.17-3.81-3.84-3.81zM5 7H4v1H3V7H2V6h1V5h1v1h1v1zm5.06 1.11a1.06 1.06 0 1 1 .001-2.121a1.06 1.06 0 0 1-.001 2.121zM13 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}