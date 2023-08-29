package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Relayreddit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.46 13.977a5.084 5.084 0 1 1 0 10.168m0-10.168H4.59m33.87 10.168H9.657m33.901 9.832s-11.172-9.828-19.855-9.828M4.558 33.966V29.23a5.091 5.091 0 0 1 5.1-5.084"/>`),
		g.Group(children),
	)
}