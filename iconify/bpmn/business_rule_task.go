package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessRuleTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g transform="translate(0 947.638)"><rect width="1800" height="1460" x="100" y="-677.638" fill="transparent" stroke="currentColor" stroke-linecap="round" stroke-width="120" rx="329.651" ry="328.5"/><path fill="currentColor" d="M404.762-425.075V-209.9h921.642v-215.175z"/><path fill="currentColor" d="M378.773-452.176v235.457h969.73v-235.457z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="43.105" d="M379.284-93.79v-358.386H1348.4v716.774H617.528v-486.713v486.713H379.284V17.32H1348.4H379.284z"/></g>`),
		g.Group(children),
	)
}