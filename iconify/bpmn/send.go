package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Send(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" stroke="currentColor"><path d="M198 484h1604l-802 454z"/><path d="m200 652l800 448l804-448v896H200z"/></g>`),
		g.Group(children),
	)
}