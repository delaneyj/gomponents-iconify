package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyPlusSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#405866" d="M58.18 25.98H38.03V5.83c0-7.77-12.05-7.77-12.05 0v20.15H5.83c-7.768 0-7.768 12.05 0 12.05h20.15v20.15c0 7.77 12.05 7.77 12.05 0V38.03h20.15c7.769 0 7.769-12.05 0-12.05"/>`),
		g.Group(children),
	)
}