package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M42 36V20H14V36C14 39.3137 16.6863 42 20 42H36C39.3137 42 42 39.3137 42 36Z"/><path d="M4 20L44 20"/><path d="M18 8V12"/><path d="M28 6V12"/><path d="M38 8V12"/></g>`),
		g.Group(children),
	)
}