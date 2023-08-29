package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneplusStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.733 25.587v-8.165m-4.398 3.766H36.5m-3.767 7.126v10.108H15.5V21.188h10.108m-1.103 13.09v-8.691"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.505 25.587h0a2.515 2.515 0 0 1-2.515 2.515h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.944 10.338c0 2.719 4.054 4.923 9.056 4.923s9.056-2.204 9.056-4.923"/><circle cx="14.944" cy="10.338" r="1.525" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.056" cy="10.338" r="1.525" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}