package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopPlusOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.828 5.828a6 6 0 0 0 6.865 9.642c.058.091.127.178.207.258l3.535 3.535a1.5 1.5 0 0 0 2.121-2.12l-3.535-3.536l-.06-.057a6.002 6.002 0 0 0-9.133-7.722Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M5.182 5.182a5.5 5.5 0 1 0 7.778 7.778a5.5 5.5 0 0 0-7.778-7.778Zm7.071 7.071A4.5 4.5 0 1 1 5.89 5.889a4.5 4.5 0 0 1 6.364 6.364Z" clip-rule="evenodd"/><path d="M12 14.121a1 1 0 0 1 1.414-1.414l3.789 3.789a1 1 0 0 1-1.414 1.414L12 14.12ZM7.207 9.664a.5.5 0 1 1 0-1h4a.5.5 0 1 1 0 1h-4Z"/><path d="M8.707 7.164a.5.5 0 0 1 1 0v4a.5.5 0 0 1-1 0v-4Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}