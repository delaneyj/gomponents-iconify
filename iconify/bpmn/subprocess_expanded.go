package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubprocessExpanded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g transform="translate(0 995.638)"><path fill="currentColor" d="M677.543 49.587V742.5h692.914V49.587H677.543zm63.468 63.489h565.978v565.956H741.01V113.076z"/><path fill="currentColor" d="M816.126 337.803v96.207h415.748v-96.207z"/><rect width="1699.302" height="1400.778" x="174.349" y="-672.027" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="100" rx="266.951"/></g>`),
		g.Group(children),
	)
}