package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoxTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="62" stroke-dashoffset="62" d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/><animate fill="freeze" attributeName="fill-opacity" begin="1.3s" dur="0.15s" values="0;0.3"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M8 8h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M8 12h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M8 16h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path></g>`),
		g.Group(children),
	)
}