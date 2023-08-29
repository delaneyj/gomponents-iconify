package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiStrengthFourAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3C7.8 3 3.7 4.4.4 7c4 5.1 7.5 9.4 11.6 14.5c2.4-3 4.7-5.8 7-8.7V8h3.8c.2-.3.6-.7.8-1c-3.3-2.6-7.4-4-11.6-4m9 7v6h2v-6m-2 8v2h2v-2"/>`),
		g.Group(children),
	)
}