package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.5 10.22h-35a2 2 0 0 0-2 2v23.3a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2v-23.3a2 2 0 0 0-2-2ZM9.38 15.44h12.26M9.38 21.11h12.26M9.38 26.79h12.26m0 5.68H9.38m16.98-17.03h12.27m-12.27 5.67h12.27m-12.27 5.68h12.27m0 5.68H26.36"/>`),
		g.Group(children),
	)
}