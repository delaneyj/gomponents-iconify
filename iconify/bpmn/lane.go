package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M96 484v1080h1856V484zm96 989.242v-891.54l1660-.003v891.543z"/>`),
		g.Group(children),
	)
}