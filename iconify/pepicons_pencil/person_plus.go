package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.95 6.75a.5.5 0 0 1 .5-.5h5.25a.5.5 0 1 1 0 1h-5.25a.5.5 0 0 1-.5-.5Z"/><path d="M15 10a.5.5 0 0 1-.5-.5V4.25a.5.5 0 0 1 1 0V9.5a.5.5 0 0 1-.5.5ZM6.8 3a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM3.3 5.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z"/><path d="M.8 14.5c0-3.322 2.67-6.5 6-6.5s6 3.178 6 6.5v2a.5.5 0 0 1-1 0v-2c0-2.873-2.32-5.5-5-5.5s-5 2.627-5 5.5v2a.5.5 0 0 1-1 0v-2Z"/></g>`),
		g.Group(children),
	)
}