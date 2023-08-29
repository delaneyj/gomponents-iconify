package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moneytracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="19.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.135 31.36a5.484 5.484 0 0 0 4.6 2.07h2.76a4.6 4.6 0 0 0 0-9.2h-2.99a4.6 4.6 0 1 1 0-9.2h2.76c2.07 0 3.45.46 4.6 2.07m-5.98-4.6v23"/>`),
		g.Group(children),
	)
}