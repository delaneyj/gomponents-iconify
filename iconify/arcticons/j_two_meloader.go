package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JTwoMeloader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40 5.5H7a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.5 9v19.77h9.89M12.61 9h9.89v14.83a5 5 0 0 1-4.95 4.94h0a4.94 4.94 0 0 1-4.94-4.94v-1.64M30.1 38.98h3.61m-3.61-7.21h3.61m-3.61 3.61h2.35m-2.35-3.61v7.21m-9.17.01v-7.2l3.6 7.21l3.61-7.2V39m-13.85-4.82a2.39 2.39 0 0 1 4.78 0a2.24 2.24 0 0 1-.7 1.69c-1 .85-4.08 3.13-4.08 3.13h4.78"/>`),
		g.Group(children),
	)
}