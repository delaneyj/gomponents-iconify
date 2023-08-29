package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyDollarSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4d5357" d="M48.9 28.2H15.1c-3.1 0-5.6-2.5-5.6-5.6c0-3.1 2.5-5.6 5.6-5.6h33.8c3.1 0 5.6 2.5 5.6 5.6H62c0-7.2-5.9-13.1-13.1-13.1H35.8V2h-7.5v7.5H15.1C7.9 9.5 2 15.4 2 22.6s5.9 13.1 13.1 13.1h33.8c3.1 0 5.6 2.5 5.6 5.6c0 3.1-2.5 5.6-5.6 5.6H15.1c-3.1 0-5.6-2.5-5.6-5.6H2c0 7.2 5.9 13.1 13.1 13.1h13.1V62h7.5v-7.5h13.1c7.2 0 13.1-5.9 13.1-13.1s-5.8-13.2-13-13.2"/>`),
		g.Group(children),
	)
}