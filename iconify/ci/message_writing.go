package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageWriting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9H7v2h2V9Zm2 0h2v2h-2V9Zm6 0h-2v2h2V9Z"/><path fill="currentColor" d="M3 5v16l4.8-3.6c.346-.26.767-.4 1.2-.4h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2Zm2 12V5h14v10H8.334a1.984 1.984 0 0 0-1.2.4L5 17Z"/>`),
		g.Group(children),
	)
}