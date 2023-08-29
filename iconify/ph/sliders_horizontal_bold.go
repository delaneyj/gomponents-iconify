package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M40 92h30.06a36 36 0 0 0 67.88 0H216a12 12 0 0 0 0-24h-78.06a36 36 0 0 0-67.88 0H40a12 12 0 0 0 0 24Zm64-24a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm112 96h-14.06a36 36 0 0 0-67.88 0H40a12 12 0 0 0 0 24h94.06a36 36 0 0 0 67.88 0H216a12 12 0 0 0 0-24Zm-48 24a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}