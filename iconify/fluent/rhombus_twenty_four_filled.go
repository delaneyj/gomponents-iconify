package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhombusTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.807 5.408A2.25 2.25 0 0 1 7.894 4h12.858a2.25 2.25 0 0 1 2.087 3.092l-4.642 11.5A2.25 2.25 0 0 1 16.111 20H3.252a2.25 2.25 0 0 1-2.086-3.092l4.641-11.5Z"/>`),
		g.Group(children),
	)
}