package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1403 39q17 41-14 70L896 602v742q0 42-39 59q-13 5-25 5q-27 0-45-19l-256-256q-19-19-19-45V602L19 109Q-12 80 5 39Q22 0 64 0h1280q42 0 59 39z"/>`),
		g.Group(children),
	)
}