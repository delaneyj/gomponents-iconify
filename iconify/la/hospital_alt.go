package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4.906L3.625 10.063l-.625.28V27h26V10.344l-.625-.281zm0 2.188l11 4.593V25H5V11.687zM15 13v3h-3v2h3v3h2v-3h3v-2h-3v-3z"/>`),
		g.Group(children),
	)
}