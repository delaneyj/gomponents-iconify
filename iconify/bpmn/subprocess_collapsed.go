package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubprocessCollapsed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g transform="translate(0 995.638)"><path fill="currentColor" d="M677.543 48.007V740.92h692.914V48.007H677.543zm63.473 63.472h565.968v565.97H741.016v-565.97zm235.505 80.502v157.701H818.817v94.963H976.52v157.702h94.958V444.645h157.704v-94.963h-157.703v-157.7h-94.957z"/><rect width="1699.302" height="1400.778" x="174.349" y="-672.027" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="100" rx="266.951"/></g>`),
		g.Group(children),
	)
}