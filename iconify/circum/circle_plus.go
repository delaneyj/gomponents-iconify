package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 12.5h-2.5V15a.5.5 0 0 1-1 0v-2.5H9a.5.5 0 0 1 0-1h2.5V9a.5.5 0 0 1 1 0v2.5H15a.5.5 0 0 1 0 1Z"/><path fill="currentColor" d="M12 21.932A9.934 9.934 0 1 1 21.932 12A9.944 9.944 0 0 1 12 21.932Zm0-18.867A8.934 8.934 0 1 0 20.932 12A8.944 8.944 0 0 0 12 3.065Z"/>`),
		g.Group(children),
	)
}