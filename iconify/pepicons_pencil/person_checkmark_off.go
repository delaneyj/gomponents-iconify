package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCheckmarkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M6.8 3a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM3.3 5.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z"/><path d="M.8 14.5c0-3.322 2.67-6.5 6-6.5s6 3.178 6 6.5v2a.5.5 0 0 1-1 0v-2c0-2.873-2.32-5.5-5-5.5s-5 2.627-5 5.5v2a.5.5 0 0 1-1 0v-2ZM18.154 3.563a.5.5 0 0 1 .194.68l-2.778 5a.5.5 0 0 1-.874-.486l2.778-5a.5.5 0 0 1 .68-.194Z"/><path d="M11.965 6.465a.5.5 0 0 1 .703-.078l2.778 2.223a.5.5 0 1 1-.625.78l-2.778-2.222a.5.5 0 0 1-.078-.703Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}