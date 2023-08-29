package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ornito(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 8.617a5.776 5.776 0 0 1 5.392-2.683c3.615 0 3.68 6.95 5.672 9.148s5.353 2.261 9.71 7.116a120.206 120.206 0 0 0 11.882 10.834l-3.398-.587l7.742 6.158s-2.35 2.849-2.632 3.245s-1.84.268-2.632-.511s-7.385-6.72-7.385-6.72s-4.273-.14-5.918-1.099l-8.557 5.851l-2.555-3.22l3.807-4.369s-5.877-4.165-6.516-9.416a35.714 35.714 0 0 1 .473-10.413A3.36 3.36 0 0 0 7.697 8.63A16.445 16.445 0 0 0 5.5 8.617Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.595 26.525l1.914-2.193l-3.143-2.913l-2.201 1.743M16.628 31.78c1.584.562 5.718 1.367 7.305 1.738"/>`),
		g.Group(children),
	)
}