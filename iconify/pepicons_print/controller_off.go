package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M6.5 6.5h9A4.5 4.5 0 0 1 20 11v2a4.5 4.5 0 0 1-4.5 4.5h-9A4.5 4.5 0 0 1 2 13v-2a4.5 4.5 0 0 1 4.5-4.5Z" opacity=".2"/><path d="M12.25 10a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Zm2 2.5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/><path fill-rule="evenodd" d="M14.5 4.5h-9A4.5 4.5 0 0 0 1 9v2a4.5 4.5 0 0 0 4.5 4.5h9A4.5 4.5 0 0 0 19 11V9a4.5 4.5 0 0 0-4.5-4.5ZM2 9a3.5 3.5 0 0 1 3.5-3.5h9A3.5 3.5 0 0 1 18 9v2a3.5 3.5 0 0 1-3.5 3.5h-9A3.5 3.5 0 0 1 2 11V9Z" clip-rule="evenodd"/><path d="M4 11a1 1 0 1 1 0-2h4a1 1 0 0 1 0 2H4Z"/><path d="M7 12a1 1 0 1 1-2 0V8a1 1 0 0 1 2 0v4Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}