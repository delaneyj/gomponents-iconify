package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificatethree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m485 594l91 366l-128 64l-86-387q-23 3-42 3t-42-3l-86 387l-128-64l91-366Q84 551 42 478.5T0 320q0-87 43-160.5T159.5 43T320 0t160.5 43T597 159.5T640 320q0 86-42 158.5T485 594z"/>`),
		g.Group(children),
	)
}