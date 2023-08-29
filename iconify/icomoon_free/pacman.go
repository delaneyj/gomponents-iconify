package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pacman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.074 2.794a8 8 0 1 0 0 10.412L10 8l5.074-5.206zM11 1.884a1.116 1.116 0 1 1 0 2.232a1.116 1.116 0 0 1 0-2.232z"/>`),
		g.Group(children),
	)
}