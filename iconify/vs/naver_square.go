package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NaverSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M64 0Q39 0 19.5 19T0 64v1664q0 26 19.5 45t44.5 19h1664q25 0 44.5-19t19.5-45V64q0-26-19.5-45T1728 0H64zm337 448h345l300 448V448h345v896h-345L746 896v448H401V448z"/>`),
		g.Group(children),
	)
}