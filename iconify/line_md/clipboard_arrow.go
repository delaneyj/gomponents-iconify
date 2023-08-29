package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M12 3H19V11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M19 17V21H5V3H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="44;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 14H12.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 14L15 17M12 14L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		g.Group(children),
	)
}