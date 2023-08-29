package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Max(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.5 19.619l-6.6 8.762m6.6 0l-6.6-8.762M7.5 28.381v-5.234c0-1.821 1.48-3.3 3.3-3.3h0c1.82 0 3.3 1.479 3.3 3.3v5.234"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.1 23.146c0-1.82 1.48-3.3 3.3-3.3h0c1.82 0 3.3 1.48 3.3 3.3v5.235m9.786-3.3c0 1.82-1.48 3.3-3.3 3.3h0c-1.82 0-3.3-1.48-3.3-3.3h0v-2.162c0-1.82 1.48-3.3 3.3-3.3h0c1.82 0 3.3 1.48 3.3 3.3m0 5.462v-8.762M27.186 25v-2"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}