package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditKarma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.133 7.636V34.37m0-5.681l12.03-12.03m-8.354 8.354L39.5 34.37m-18.635-3.345a6.468 6.468 0 0 1-5.681 3.342h0A6.703 6.703 0 0 1 8.5 27.683V23.34a6.703 6.703 0 0 1 6.684-6.684h0a6.468 6.468 0 0 1 5.68 3.342"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}