package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hobi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-30 13.506V29.59m7.012-10.584V29.59M10.5 24.278h7.012"/><rect width="5.292" height="7.012" x="20.333" y="22.578" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.646" ry="2.646"/><circle cx="36.574" cy="19.337" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.574 22.578v7.012m-8.111-4.366a2.646 2.646 0 0 1 2.646-2.646h0a2.646 2.646 0 0 1 2.646 2.646v1.72a2.646 2.646 0 0 1-2.646 2.646h0a2.646 2.646 0 0 1-2.646-2.646m0 2.646V19.006"/>`),
		g.Group(children),
	)
}