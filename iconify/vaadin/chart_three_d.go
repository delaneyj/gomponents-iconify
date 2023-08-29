package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 4V2L8 0L4 2v1L0 5v5l12 6l4-2V6zm-8 6.88l-3-1.5v-3.3l3 1.53v3.27zm0-4.39l-2.34-1.2L4 4.12v2.37zm4 6.39l-3-1.5V3.07l3 1.54v8.27zM5.66 2.29L8 1.12l2.34 1.17L8 3.49zM12 14.88l-3-1.5V7.07l3 1.54v6.27zm0-7.39l-2.34-1.2L12 5.12l2.34 1.17z"/>`),
		g.Group(children),
	)
}