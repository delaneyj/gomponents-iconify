package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M429.65 210C214.934 210 40 384.196 40 598.5v803c0 214.304 174.934 388.5 389.65 388.5h1140.7c214.716 0 389.65-174.196 389.65-388.5v-803c0-214.304-174.934-388.5-389.65-388.5H429.65zm0 120h1140.7C1720.887 330 1840 448.826 1840 598.5v803c0 149.674-119.113 268.5-269.65 268.5H429.65C279.113 1670 160 1551.174 160 1401.5v-803C160 448.826 279.113 330 429.65 330z"/>`),
		g.Group(children),
	)
}