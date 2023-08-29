package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.879 13.536L8.464 12.12L10.586 10L8.464 7.879L9.88 6.464L12 8.586l2.121-2.122l1.415 1.415L13.414 10l2.122 2.121l-1.415 1.415L12 11.414l-2.121 2.122Z"/><path fill="currentColor" d="M3 5v16l4.8-3.6c.346-.26.767-.4 1.2-.4h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2Zm2 12V5h14v10H8.334a1.984 1.984 0 0 0-1.2.4L5 17Z"/>`),
		g.Group(children),
	)
}