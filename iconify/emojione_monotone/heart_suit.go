package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M46.063 2c-6.268 0-11.515 3.598-14.062 8.81C29.452 5.598 24.206 2 17.938 2C9.227 2 2 9.361 2 17.938C2 32.406 32.001 62 32.001 62S62 32.406 62 17.938C62 9.361 54.775 2 46.063 2z"/>`),
		g.Group(children),
	)
}