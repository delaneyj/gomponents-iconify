package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCircleTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g stroke-dasharray="12" stroke-dashoffset="12"><path d="M12 7V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path></g>`),
		g.Group(children),
	)
}