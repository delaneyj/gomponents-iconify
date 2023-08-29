package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 30.4L96 168.681V232h320v-63.319ZM384 200H128v-16.681L256 72.7l128 110.619ZM96 343.319L256 481.6l160-138.281V280H96ZM128 312h256v16.681L256 439.3L128 328.681Z"/>`),
		g.Group(children),
	)
}