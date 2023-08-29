package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Totalcommander(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95s-3.29.1-5.5.11a1.94 1.94 0 0 0-2 1.94v9.23a1 1 0 0 1-1 1H15.87a1 1 0 0 1-1-1V7.45A2 2 0 0 0 13 5.5ZM27.29 7a1 1 0 0 0-1 1v6.56a1 1 0 0 0 1 1h1.39a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1ZM12.54 21.92h22.92a2 2 0 0 1 1.95 1.95v14a1.94 1.94 0 0 1-1.95 1.94H12.54a1.94 1.94 0 0 1-1.95-1.94v-14a2 2 0 0 1 1.95-1.95Zm.66 4.58h21.6m0 4.36H13.2m0 4.35h21.6"/>`),
		g.Group(children),
	)
}