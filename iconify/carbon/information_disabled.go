package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 3.414L28.586 2L2 28.586L3.414 30l3.443-3.443a13.961 13.961 0 0 0 19.7-19.7zM28 16a11.973 11.973 0 0 1-19.732 9.146L15 18.414V22h-2v2h7v-2h-3v-5.586l8.146-8.146A11.897 11.897 0 0 1 28 16zM16 8a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 16 8z"/><path fill="currentColor" d="M5.67 22.085A11.983 11.983 0 0 1 22.086 5.67l1.454-1.454A13.985 13.985 0 0 0 4.216 23.54Z"/>`),
		g.Group(children),
	)
}