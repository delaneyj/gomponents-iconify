package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListIndented(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 5H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.2s" values="14;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 10H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 20H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><g fill="currentColor" fill-opacity="0"><circle cx="4" cy="5" r="1"><animate fill="freeze" attributeName="fill-opacity" dur="0.2s" values="0;1"/></circle><circle cx="6" cy="10" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.2s" values="0;1"/></circle><circle cx="8" cy="15" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></circle><circle cx="10" cy="20" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></g>`),
		g.Group(children),
	)
}