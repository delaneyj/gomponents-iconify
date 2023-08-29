package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mifareclassictool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.38 18.62a2.48 2.48 0 0 1 2.46 2.47h0m-2.46-4.37A4.27 4.27 0 0 1 37.67 21h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.38 14.89A6.12 6.12 0 0 1 39.5 21h0m-31 8.51v-8.58l4.29 8.58l4.3-8.58v8.58m10.01-8.58h5.72m-2.86 8.58v-8.58m-4.61 5.72h0a2.85 2.85 0 0 1-2.86 2.86h0a2.85 2.85 0 0 1-2.86-2.86v-2.86a2.85 2.85 0 0 1 2.86-2.86h0a2.85 2.85 0 0 1 2.86 2.86h0"/><rect width="39" height="31.2" x="4.5" y="8.4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.95"/>`),
		g.Group(children),
	)
}