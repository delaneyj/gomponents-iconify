package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 48v416h416V48Zm170 312.38l-80.6-89.57l23.79-21.41l56 62.22L350 153.46L374.54 174Z"/>`),
		g.Group(children),
	)
}