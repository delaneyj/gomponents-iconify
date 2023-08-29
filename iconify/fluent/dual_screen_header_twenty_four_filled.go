package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenHeaderTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.995 7.997V18.25a1.75 1.75 0 0 1-1.75 1.75h-7.247a1.78 1.78 0 0 1-.255-.019l-.001-11.984h9.253Zm-10.752 0V19.98a1.8 1.8 0 0 1-.245.017H3.75A1.75 1.75 0 0 1 2 18.247V7.997h9.243Zm-.245-3.995c.083 0 .165.006.246.017l-.001 2.978H2V5.752c0-.967.784-1.75 1.75-1.75h7.248Zm9.248.002c.966 0 1.75.784 1.75 1.75l-.001 1.243h-9.253V4.023c.084-.012.17-.019.256-.019h7.248Z"/>`),
		g.Group(children),
	)
}