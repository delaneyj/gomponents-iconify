package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.67 9.69L14 5.84a2.67 2.67 0 0 0-4 2.31L6 5.84a2.67 2.67 0 0 0-4 2.31v7.7a2.63 2.63 0 0 0 1.33 2.3a2.61 2.61 0 0 0 1.34.37A2.69 2.69 0 0 0 6 18.16l4-2.31a2.65 2.65 0 0 0 1.33 2.31a2.66 2.66 0 0 0 2.67 0l6.67-3.85a2.67 2.67 0 0 0 0-4.62ZM10 13.54l-5 2.88a.67.67 0 0 1-1-.57v-7.7a.67.67 0 0 1 1-.57l5 2.88Zm9.67-1L13 16.43a.69.69 0 0 1-.67 0a.66.66 0 0 1-.33-.58v-7.7a.66.66 0 0 1 .33-.58a.78.78 0 0 1 .34-.09a.63.63 0 0 1 .33.09l6.67 3.85a.67.67 0 0 1 0 1.16Z"/>`),
		g.Group(children),
	)
}