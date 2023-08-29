package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentValley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.313 15.78a4.703 4.703 0 0 0-1.62-6.45a5.694 5.694 0 0 0-1.395-.558C26.41 8.235 12.018 4.5 12.018 4.5s10.465 10.854 12.079 12.332a4.708 4.708 0 0 0 7.216-1.051Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.869 9.438a4.705 4.705 0 0 0-1.067 8.338"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.299 8.772c-3.565 1.508-4.519 4.546-4.202 8.06m1.871 1.049c0 4.634-8.916 7.728-8.916 10.716s2.916 5.833 8.038 5.833s6.651-2.276 6.651-3.948s-3.443-6.952-3.443-12.528"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.463 19.791a3.486 3.486 0 0 0 2.994.307m-3.189 14.331v7.648l1.46 1.423m1.317-9.388v6.054l1.286 1.12"/>`),
		g.Group(children),
	)
}