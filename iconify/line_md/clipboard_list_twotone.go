package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardListTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 10H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		g.Group(children),
	)
}