package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelHogwarts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m42 501l121-106q4-5 15-5q5 0 15 5l122 106q4 5 15 5q8 0 14.5-6.5T351 484V21q0-9-6-15t-15-6H27Q17 0 11 6T5 21v465q0 9 6.5 15.5T27 508q8 0 15-7zM261 43h43v396l-43-38V43zm-128 0h86v322q-15-15-43-15q-21 0-43 15V43zm-85 0h43v358l-43 38V43z"/>`),
		g.Group(children),
	)
}