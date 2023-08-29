package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thunder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.658 40.755c2.631-1.655 4.418-7.282 1.028-10.234L9.773 42.124c-.64.936-.3 1.934 1.695 1.005l7.296-3.405l12.395-5.77a3.234 3.234 0 0 0 1.502-4.32l-.006-.012l-6.84-14.775c-.82-1.689.484-4.355 2.886-6.314m2.64-1.288c-2.63 1.655-4.418 7.282-1.027 10.234l7.913-11.603c.64-.936.3-1.934-1.695-1.005l-7.296 3.405l-12.395 5.77a3.234 3.234 0 0 0-1.502 4.32l.006.012l6.84 14.775c.82 1.689-.484 4.355-2.886 6.315"/>`),
		g.Group(children),
	)
}