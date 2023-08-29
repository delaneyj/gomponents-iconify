package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telegram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M21 5L18.5 20M21 5L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="22" stroke-dashoffset="22" d="M21 5L2 12.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="22;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M18.5 20L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M2 12.5L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 16L9 19M9 13.5L9 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="6;0"/></path></g>`),
		g.Group(children),
	)
}