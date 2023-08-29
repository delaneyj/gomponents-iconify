package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gsa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.337 25.77h-4.631m-1.156 3.544l3.475-10.66L38.5 29.346m-17.763-1.172c.656.854 1.478 1.172 2.621 1.172h1.583a2.667 2.667 0 0 0 2.667-2.667v-.012A2.667 2.667 0 0 0 24.941 24h-1.746a2.67 2.67 0 0 1-2.67-2.67h0a2.676 2.676 0 0 1 2.676-2.676h1.574c1.144 0 1.966.318 2.622 1.172m-10.814 2.37a3.542 3.542 0 0 0-3.542-3.542h0A3.542 3.542 0 0 0 9.5 22.196v3.608a3.542 3.542 0 0 0 3.541 3.542h0a3.542 3.542 0 0 0 3.542-3.542H13.04"/>`),
		g.Group(children),
	)
}