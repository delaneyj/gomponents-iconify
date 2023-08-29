package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g stroke="currentColor" transform="translate(0 947.638)"><rect width="1800" height="1460" x="100" y="-677.638" fill="transparent" stroke-linecap="round" stroke-width="120" rx="329.651" ry="328.5"/><g fill="currentColor" fill-rule="evenodd" stroke-width=".623"><path d="M346.858-428.042h999.853l-499.927 283z"/><path d="m348.104-323.32l498.68 279.261l501.174-279.26V235.2H348.104z"/></g></g>`),
		g.Group(children),
	)
}