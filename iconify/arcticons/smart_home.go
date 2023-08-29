package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.756 27.073v6.83a1.043 1.043 0 0 0 1.043 1.044h4.715a1.043 1.043 0 0 0 1.044-1.044v-8.515h1.368a1.571 1.571 0 0 0 1.018-2.767l-10.698-9.11a1.922 1.922 0 0 0-2.492 0l-10.698 9.11a1.57 1.57 0 0 0 1.018 2.767h1.368v8.516a1.044 1.044 0 0 0 1.044 1.043H20.2a1.043 1.043 0 0 0 1.043-1.043v-6.83Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.42 5.5H7.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2V19.58"/>`),
		g.Group(children),
	)
}