package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 18H4V42H44V18H38"/><path fill="#2F88FF" d="M38 6H10V22.5L24 33L38 22.5V6Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 22.5L24 33L38 22.5M10 22.5V6H38V22.5M10 22.5L4 18M38 22.5L44 18"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 19L24 24M24 24L29 19M24 24V13"/></g>`),
		g.Group(children),
	)
}