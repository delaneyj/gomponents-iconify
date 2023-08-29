package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 24v2h-6v-2z"/><path fill="currentColor" d="M25 30H7a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h18a2.002 2.002 0 0 1 2 2v24a2.003 2.003 0 0 1-2 2ZM7 4v24h18V4Z"/>`),
		g.Group(children),
	)
}