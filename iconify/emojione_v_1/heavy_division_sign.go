package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyDivisionSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#405866" d="M58.02 26.556H6.135c-7.699 0-7.699 11.94 0 11.94H58.02c7.701 0 7.701-11.94 0-11.94M34.899 6.203h-5.948c-7.699 0-7.699 13.12 0 13.12h5.948c7.701 0 7.701-13.12 0-13.12m.161 39.68h-5.949c-7.698 0-7.698 13.12 0 13.12h5.949c7.7 0 7.7-13.12 0-13.12"/>`),
		g.Group(children),
	)
}