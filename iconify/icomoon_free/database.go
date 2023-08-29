package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0C3.582 0 0 1.119 0 2.5v2C0 5.881 3.582 7 8 7s8-1.119 8-2.5v-2C16 1.119 12.418 0 8 0z"/><path fill="currentColor" d="M8 8.5C3.582 8.5 0 7.381 0 6v3c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5V6c0 1.381-3.582 2.5-8 2.5z"/><path fill="currentColor" d="M8 13c-4.418 0-8-1.119-8-2.5v3C0 14.881 3.582 16 8 16s8-1.119 8-2.5v-3c0 1.381-3.582 2.5-8 2.5z"/>`),
		g.Group(children),
	)
}