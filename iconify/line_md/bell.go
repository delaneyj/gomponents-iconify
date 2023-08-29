package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		g.Group(children),
	)
}