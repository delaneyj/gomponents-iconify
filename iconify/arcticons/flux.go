package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.51 21.51 0 0 0 2.5 24a22 22 0 0 0 .24 3a21.48 21.48 0 0 0 39.43 8.46s0 0 0 0c.37-.6.72-1.21 1-1.85c0 0 0 0 0 0a17.73 17.73 0 0 0 .83-1.89a3.81 3.81 0 0 0 .15-.44c.2-.53.37-1.06.52-1.61c0-.16.08-.32.11-.48c.14-.52.25-1 .35-1.58c0-.3.08-.59.11-.89s.1-.83.14-1.25s0-.95 0-1.43A21.51 21.51 0 0 0 24 2.5ZM2.74 27a20.6 20.6 0 0 1 4.53-4.94a15.16 15.16 0 0 1 9.3-3.24h0c3.35.06 6.55 1.39 9.56 2.89c.36.17.72.37 1.08.55a6.43 6.43 0 1 1 10.54 7.34a34.72 34.72 0 0 1 4.56 5.65m-15.1-13a56.69 56.69 0 0 1 8.65 5.57c.33.28 1.57 1.46 1.89 1.74m-5.06-9.13"/>`),
		g.Group(children),
	)
}