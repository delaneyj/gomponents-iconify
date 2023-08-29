package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessRule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M212.277 418v382.37h1574.785V418z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="70" d="M213.105 1000V418h1573.79v1164H600V791.606V1582H213.105v-401.564h1573.79h-1573.79z"/>`),
		g.Group(children),
	)
}