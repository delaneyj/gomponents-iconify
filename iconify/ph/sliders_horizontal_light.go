package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M40 86h34.6a30 30 0 0 0 58.8 0H216a6 6 0 0 0 0-12h-82.6a30 30 0 0 0-58.8 0H40a6 6 0 0 0 0 12Zm64-24a18 18 0 1 1-18 18a18 18 0 0 1 18-18Zm112 108h-18.6a30 30 0 0 0-58.8 0H40a6 6 0 0 0 0 12h98.6a30 30 0 0 0 58.8 0H216a6 6 0 0 0 0-12Zm-48 24a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}