package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeArrowFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M14 4V11C14 11.7956 13.6839 12.5587 13.1213 13.1213C12.5587 13.6839 11.7956 14 11 14H7C6.20435 14 5.44129 13.6839 4.87868 13.1213C4.31607 12.5587 4 11.7956 4 11V4z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.5s" values="0;1"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M14 9H17C17.55 9 18 8.55 18 8V5C18 4.45 17.55 4 17 4H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 18H19.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.3s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M19.5 18l-3 -3M19.5 18l-3 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g>`),
		g.Group(children),
	)
}