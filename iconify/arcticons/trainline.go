package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trainline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95Z"/><rect width="4.79" height="6.37" x="18.69" y="31.78" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.48 31.78v6.38m4.78-3.98a2.4 2.4 0 0 1 2.4-2.4h0a2.4 2.4 0 0 1 2.39 2.4v4m-4.79-6.4v6.38m-13.77-3.98a2.4 2.4 0 0 1 2.4-2.4h0m-2.4 0v6.38m-3.85-9.59v9.59m-1.68-6.38h3.35m13.56 0v6.38"/><circle cx="25.87" cy="29.31" r=".9" fill="currentColor"/>`),
		g.Group(children),
	)
}