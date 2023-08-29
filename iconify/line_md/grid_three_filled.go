package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridThreeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round"><g><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></g><g><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.5s" values="0;1"/></g><g><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.5s" values="0;1"/></g></g>`),
		g.Group(children),
	)
}