package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitpay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.144 25.97a7.009 7.009 0 0 1 6.624-5.63h0a4.595 4.595 0 0 1 4.638 5.63l-.645 3.66a7.01 7.01 0 0 1-6.624 5.632h0a4.595 4.595 0 0 1-4.638-5.631m-.993 5.631l3.972-22.524"/>`),
		g.Group(children),
	)
}