package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m148 739l241-125l241 125l-46-267l194-189l-269-39L389 0L269 244L0 283l195 189z"/>`),
		g.Group(children),
	)
}