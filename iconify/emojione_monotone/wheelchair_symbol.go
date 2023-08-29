package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-1.001 7a4 4 0 1 1 0 8A4 4 0 0 1 27 13c0-2.211 1.79-4 3.999-4zm-1.624 44C22.552 53 17 47.459 17 40.646c0-5.369 3.45-9.947 8.25-11.646v4.518a8.235 8.235 0 0 0-4.123 7.129c0 4.541 3.699 8.236 8.248 8.236a8.254 8.254 0 0 0 7.166-4.168L39 48.396A12.36 12.36 0 0 1 29.375 53zm16.107-1l-7.373-11.408H29V20h3.991v8.236h7.982v4.119h-7.982v4.117h7.253l5.932 9.182l2.48-2.656L51 46.402L45.482 52z"/>`),
		g.Group(children),
	)
}