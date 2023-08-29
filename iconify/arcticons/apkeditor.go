package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apkeditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.078 14.34a1.37 1.37 0 1 1 1.36-1.316a1.37 1.37 0 0 1-1.36 1.315Zm7.934 0a1.37 1.37 0 1 1 1.359-1.316a1.37 1.37 0 0 1-1.36 1.315Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.261 17.894v-2.229a9.717 9.717 0 0 1 9.762-9.716h0a9.716 9.716 0 0 1 9.716 9.716v2.229M15.643 4.5l2.451 3.009M32.424 4.5l-2.518 3.009"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.236 15.086V35.7a3.134 3.134 0 0 0 2.786 2.786h22.843V23.443h1.114a3.134 3.134 0 0 0 2.786-2.786a3.134 3.134 0 0 0-2.786-2.786H12.022a3.134 3.134 0 0 1-2.786-2.785a3.134 3.134 0 0 1 2.786-2.786h2.73"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.864 38.486V43.5h1.115a3.134 3.134 0 0 0 2.786-2.786V20.657"/><circle cx="24" cy="26.229" r="2.229" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 23.68v-2.164m6.686 20.87L24 28.457l-6.686 13.929"/>`),
		g.Group(children),
	)
}