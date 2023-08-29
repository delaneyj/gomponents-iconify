package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-dasharray="14" stroke-dashoffset="14" stroke-linecap="round" stroke-width="2"><path d="M8 5H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.2s" values="14;0"/></path><path d="M8 10H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="14;0"/></path><path d="M8 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><path d="M8 20H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><circle cx="4" cy="5" r="1"><animate fill="freeze" attributeName="fill-opacity" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="10" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="15" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="20" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></g>`),
		g.Group(children),
	)
}