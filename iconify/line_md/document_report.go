package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-width="2"><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 17V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 17V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M15 17V12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g></g>`),
		g.Group(children),
	)
}