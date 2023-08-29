package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func News(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.373 11.5h16.502v5H20.373v-5Zm0 10h16.502v5H20.373v-5Zm0 10h16.502v5H20.373v-5Zm-9.248-19.999h5v5h-5zm0 9.999h5v5h-5zm0 10h5v5h-5z"/>`),
		g.Group(children),
	)
}