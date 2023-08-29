package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="2" d="M25.491 23.333L23.999 25H23v-1h1a2 2 0 0 0 1.491-.667ZM8.001 25l-1.492-1.667A2 2 0 0 0 8 24h1v1h-.999ZM6 5h20v10H6V5Z"/>`),
		g.Group(children),
	)
}