package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForLiberia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m27.369 35.9H4.631A27.983 27.983 0 0 1 4 32h56c0 2.024-.221 3.997-.631 5.9M10.323 49.7a28.104 28.104 0 0 1-3.707-5.9h50.768a28.06 28.06 0 0 1-3.706 5.9H10.323M32 20.2h25.384a27.762 27.762 0 0 1 1.985 5.899H32V20.2m-6.33 5.9l-5.47-3.693l-5.47 3.693l2.064-6.005l-5.445-3.721l6.745-.017l2.105-5.99l2.105 5.99l6.745.017l-5.448 3.721L25.67 26.1m28.008-11.8H32V8.4h15.039a28.234 28.234 0 0 1 6.639 5.9M16.96 55.6h30.079C42.691 58.38 37.533 60 32 60s-10.691-1.62-15.04-4.4"/>`),
		g.Group(children),
	)
}