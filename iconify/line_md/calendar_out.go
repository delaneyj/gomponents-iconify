package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" stroke-dasharray="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="0;64"/></rect><path stroke-dasharray="6" d="M7 4V2M17 4V2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="0;6"/></path><path stroke-dasharray="12" d="M7 11H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="0;12"/></path><path stroke-dasharray="9" d="M7 15H14"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="0;9"/></path></g><rect width="14" height="3" x="5" y="5" fill="currentColor"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="3;0"/></rect>`),
		g.Group(children),
	)
}