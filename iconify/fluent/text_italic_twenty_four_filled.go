package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.649 18.5h3.847a1 1 0 0 1 0 2h-9.5a.997.997 0 0 1-.996-1c0-.552.445-1 .996-1h3.51L13.332 6H9.997a1 1 0 0 1 0-2H18.5a1 1 0 0 1 0 2h-3.025L10.65 18.5Z"/>`),
		g.Group(children),
	)
}