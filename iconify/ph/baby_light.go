package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 138a10 10 0 1 1 10-10a10 10 0 0 1-10 10Zm72-20a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm-11.2 44.92a47 47 0 0 1-49.6 0a6 6 0 0 0-6.4 10.16a59 59 0 0 0 62.4 0a6 6 0 1 0-6.4-10.16ZM230 128A102 102 0 1 1 128 26a102.12 102.12 0 0 1 102 102Zm-12 0a90.11 90.11 0 0 0-87.07-89.95C118.3 55.23 118 71.85 118 72a10 10 0 0 0 20 0a6 6 0 0 1 12 0a22 22 0 0 1-44 0c0-.75.15-15.82 10.14-33.22A90 90 0 1 0 218 128Z"/>`),
		g.Group(children),
	)
}