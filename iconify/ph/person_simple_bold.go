package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 84a36 36 0 1 0-36-36a36 36 0 0 0 36 36Zm0-48a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm106.29 102.17a12 12 0 0 1-16.47 4.12c-.32-.19-32.37-18.92-77.82-21.88v27L201 216a12 12 0 1 1-18 16l-55-61.91L73 232a12 12 0 1 1-18-16l61-68.59v-27c-45.72 2.95-77.48 21.68-77.82 21.89a12 12 0 1 1-12.35-20.58C27.58 120.66 69.35 96 128 96s100.42 24.66 102.17 25.71a12 12 0 0 1 4.12 16.46Z"/>`),
		g.Group(children),
	)
}