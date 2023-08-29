package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotTickets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Zm-9.99 13.23h3.5m0 10.54V18.73"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 29.27V18.72h2.37c2.55 0 4.61 2.07 4.61 4.61v1.32c0 2.55-2.07 4.61-4.61 4.61H10.5Z"/><rect width="6.99" height="10.55" x="20.94" y="18.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.65" ry="2.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 18.73h0"/>`),
		g.Group(children),
	)
}