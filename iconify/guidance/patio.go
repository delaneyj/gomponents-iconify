package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 19.5h10m.5-5l-2-8H7.067a3 3 0 0 0-2.803 1.932L1 17m14 6.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-7v-6M15.069 0c0 .33-.05.664-.206.955A3.811 3.811 0 0 1 11.5 2.969a3.814 3.814 0 0 1 3.569 2.47a3.814 3.814 0 0 1-.778 4.27a3.814 3.814 0 0 1 4.27-.778a3.814 3.814 0 0 1 2.47 3.569c0-1.455.815-2.72 2.014-3.363c.291-.156.624-.206.955-.206M19 .526A3.177 3.177 0 1 0 23.474 5m-16.28-.5s-1.81-.557-2.134-1.776A1.768 1.768 0 0 1 6.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61l-.291.079Z"/>`),
		g.Group(children),
	)
}