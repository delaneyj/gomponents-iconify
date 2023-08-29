package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AquariumEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8 1c-.876 0-1.85.092-3.004.527C3.843 1.962 2.848 2.657 2 3.5c-.852.847-2 2.5-2 3s1.135 1.943 2.678 2.621c1.542.678 2.39.798 3.283.895c.797.086 1.942-.027 2.885-.233C9.592 9.621 10.994 9.31 11 9c0 0-2.756-.063-3-.5c-.249-.445-.25-1.586 0-2c.258-.428 2.5 1 2.5 1c.644.258.644-4.258 0-4c0 0-2.277 1.447-2.5 1c-.25-.5-.25-1.5 0-2c.223-.447 3-.5 3-.5c0-.5-2.124-1-3-1zM3.514 4.502a1.014 1.014 0 1 1 0 2.028a1.014 1.014 0 0 1 0-2.028z" fill="currentColor"/>`),
		g.Group(children),
	)
}