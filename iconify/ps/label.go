package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M325 0H27Q17 0 11 6T5 21v465q0 9 6.5 15.5T27 508q7 0 15-5l121-106q4-4 15-4q7 0 15 4l122 106q4 5 15 5q8 0 14.5-6.5T351 486V21q-7-21-26-21zm-21 439l-85-74q-15-15-43-15q-21 0-43 15l-85 74V43h256v396z"/>`),
		g.Group(children),
	)
}