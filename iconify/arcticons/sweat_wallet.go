package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SweatWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.817 9.722l11.599 20.066l-8.001 8.49h-2.992L4.5 18.03l8.399-8.308h2.918z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.538 17.89l8.445-8.168h2.918L43.5 29.788l-8.001 8.49h-2.992l-5.091-8.49"/>`),
		g.Group(children),
	)
}