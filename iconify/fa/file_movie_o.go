package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMovieO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22zm384 1528V640H992q-40 0-68-28t-28-68V128H128v1536h1280zM768 768q52 0 90 38t38 90v384q0 52-38 90t-90 38H384q-52 0-90-38t-38-90V896q0-52 38-90t90-38h384zm492 2q20 8 20 30v576q0 22-20 30q-8 2-12 2q-14 0-23-9l-265-266v-90l265-266q9-9 23-9q4 0 12 2z"/>`),
		g.Group(children),
	)
}