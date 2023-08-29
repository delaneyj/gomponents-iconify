package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.959 9.75a.75.75 0 0 1-.75-.75V2a.75.75 0 0 1 1.5 0v7a.75.75 0 0 1-.75.75Z"/><path fill-rule="evenodd" d="M7 5.75a3.75 3.75 0 1 1 0 7.5a3.75 3.75 0 0 1 0-7.5ZM9.25 9.5a2.25 2.25 0 1 0-4.5 0a2.25 2.25 0 0 0 4.5 0ZM17 18.25a3.75 3.75 0 1 0 0-7.5a3.75 3.75 0 0 0 0 7.5Zm2.25-3.75a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0Z" clip-rule="evenodd"/><path d="M6.209 15a.75.75 0 0 1 1.5 0v7a.75.75 0 0 1-1.5 0v-7Zm10.75 7.75a.75.75 0 0 1-.75-.75v-2a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-.75.75ZM6.209 2a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-1.5 0V2Z"/></g>`),
		g.Group(children),
	)
}