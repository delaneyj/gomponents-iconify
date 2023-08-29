package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sauna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15.5 1s-2.25 2.571 0 6s0 6 0 6m3-12s-2.25 2.571 0 6s0 6 0 6m3-12s-2.25 2.571 0 6s0 6 0 6M0 19.5h10m.5-5l-2-8H7.067a3 3 0 0 0-2.803 1.932L1 17m14 6.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-7v-6m1.695-6s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 6.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61l-.291.079Z"/>`),
		g.Group(children),
	)
}