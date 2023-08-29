package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><g stroke-dasharray="10" stroke-dashoffset="10"><circle cx="5" cy="5" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></circle><circle cx="5" cy="12" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></circle><circle cx="5" cy="19" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="10;0"/></circle></g><g stroke-dasharray="28" stroke-dashoffset="28"><rect width="11" height="3" x="9.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.5s" values="28;0"/></rect><rect width="11" height="3" x="9.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.5s" values="28;0"/></rect><rect width="11" height="3" x="9.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.5s" values="28;0"/></rect></g></g>`),
		g.Group(children),
	)
}