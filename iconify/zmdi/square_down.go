package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m235 288l-86-85h64V11h43v192h64zM427 11q17 0 29.5 12.5T469 53v299q0 18-12.5 30.5T427 395H43q-18 0-30.5-12.5T0 352V53q0-17 12.5-29.5T43 11h128v42H43v299h384V53H299V11h128z"/>`),
		g.Group(children),
	)
}