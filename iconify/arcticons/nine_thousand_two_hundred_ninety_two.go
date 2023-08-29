package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineThousandTwoHundredNinetyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2ZM24 5.5v37M5.5 24h37"/><ellipse cx="14.75" cy="13.06" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.45" ry="3.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.61 18.53a3.42 3.42 0 0 0 2.9 1.22h.24a3.38 3.38 0 0 0 3.45-3.31v-3.38m11.74 0a3.31 3.31 0 0 1 6.62 0a3.05 3.05 0 0 1-1 2.34c-1.34 1.18-5.65 4.35-5.65 4.35h6.62M11.44 31.56a3.31 3.31 0 0 1 6.62 0a3.05 3.05 0 0 1-1 2.34c-1.34 1.18-5.65 4.35-5.65 4.35h6.62"/><ellipse cx="33.25" cy="31.56" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.45" ry="3.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.11 37A3.42 3.42 0 0 0 33 38.25h.24a3.38 3.38 0 0 0 3.45-3.31v-3.38"/>`),
		g.Group(children),
	)
}