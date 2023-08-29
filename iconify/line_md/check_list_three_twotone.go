package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckListThreeTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g fill="none" stroke-dasharray="10" stroke-dashoffset="10" stroke-width="2"><path d="M3 5L5 7L9 3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path d="M3 12L5 14L9 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path d="M3 19L5 21L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0" stroke-dasharray="22" stroke-dashoffset="22"><rect width="9" height="3" x="11.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.15s" values="0;0.3"/></rect><rect width="9" height="3" x="11.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.9s" dur="0.15s" values="0;0.3"/></rect><rect width="9" height="3" x="11.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.15s" values="0;0.3"/></rect></g></g>`),
		g.Group(children),
	)
}