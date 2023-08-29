package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrewWrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1511.45 292.911a415.8 415.8 0 0 0-415.8 415.8a415.8 415.8 0 0 0 415.8 415.8a415.8 415.8 0 0 0 397.194-292.826l-298.378 79.95l-232.313-152.986l124.695-248.645l313.27-83.94A415.8 415.8 0 0 0 1511.45 292.91zm-492.602 555.194L96.6 1480.02c-28.476 84.174 63.924 222.774 158.995 227.068l902.881-618.354a415.8 415.8 0 0 1-139.626-240.63z"/>`),
		g.Group(children),
	)
}