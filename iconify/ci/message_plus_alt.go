package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessagePlusAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 10h-2V7h-3V5h3V2h2v3h3v2h-3v3Z"/><path fill="currentColor" d="M21 12h-2v3H8.334a1.984 1.984 0 0 0-1.2.4L5 17V5h7V3H5a2 2 0 0 0-2 2v16l4.8-3.6c.346-.26.767-.4 1.2-.4h10a2 2 0 0 0 2-2v-3Z"/>`),
		g.Group(children),
	)
}