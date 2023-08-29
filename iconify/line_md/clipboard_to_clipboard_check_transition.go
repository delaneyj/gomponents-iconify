package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardToClipboardCheckTransition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`),
		g.Group(children),
	)
}