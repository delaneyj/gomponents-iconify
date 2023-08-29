package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mongodb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M319 901q1 15 1 27q0 40-9.5 68t-22.5 28t-22.5-28t-9.5-68q0-13 2-27q-16-18-55-53t-69.5-66.5T70 705T18.5 602.5T0 480q0-68 19-131t48-107.5t65-85.5t67-67.5T255.5 38T288 0q7 14 32.5 38T377 88.5t67 67.5t65 85.5T557 349t19 131q0 65-18.5 122.5T506 705t-63.5 76.5T373 848t-54 53z"/>`),
		g.Group(children),
	)
}