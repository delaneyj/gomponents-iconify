package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stickynotealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1023 576q0 48-34 82L658 989q-34 34-82 34v1H0V0h1024v576h-1zM621 928l307-307q18-18 26-45H576v378q27-8 45-26zM960 64H64v896h448V512h448V64z"/>`),
		g.Group(children),
	)
}