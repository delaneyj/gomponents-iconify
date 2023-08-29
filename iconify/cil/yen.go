package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m396.641 97.81l-25.282-19.62l-114.946 148.122L148.938 78.587l-25.876 18.826L238.438 256h-85.967v32H240v34.823h-87.529v32H240V432h32v-77.177h87.529v-32H272V288h87.529v-32h-85.65L396.641 97.81z"/>`),
		g.Group(children),
	)
}