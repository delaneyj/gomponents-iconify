package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M318.141 73.72v167.244H0v487.372h541.219V628.314h117.562v100.021H1200V240.964H881.808V73.72H318.141zm86.456 86.456h390.755v80.788H404.597v-80.788zM0 822.188v304.092h1200V822.188H658.781v96.617H541.219v-96.617H0z"/>`),
		g.Group(children),
	)
}