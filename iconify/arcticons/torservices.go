package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torservices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.797 39.203A21.5 21.5 0 1 0 24 2.5m-6.06.866A21.512 21.512 0 0 0 5.589 35.109"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.985 24a14.985 14.985 0 1 0-4.39 10.596"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.53 24A8.47 8.47 0 1 0 24 15.53"/>`),
		g.Group(children),
	)
}