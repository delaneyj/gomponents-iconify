package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m62.397 8l819 467q16 9 16 39.5t-16 37.5l-819 467q-12 8-30 5.5t-32-17.5V22q31-34 62-14z"/>`),
		g.Group(children),
	)
}