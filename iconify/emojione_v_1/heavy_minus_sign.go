package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyMinusSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#405866" d="M58.05 26.705H5.82c-7.75 0-7.75 12.02 0 12.02h52.23c7.752 0 7.752-12.02 0-12.02"/>`),
		g.Group(children),
	)
}