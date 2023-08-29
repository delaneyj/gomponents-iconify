package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Participant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M96 402v1280h1856V402zm1756 97.7v1091.542H480V499.702l1372-.003zm-1660 0h192v1091.541H192z"/>`),
		g.Group(children),
	)
}