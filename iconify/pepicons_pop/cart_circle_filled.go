package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCartCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M19.513 9h-9.026a2.5 2.5 0 0 0-2.437 3.057l.913 4A2.5 2.5 0 0 0 11.401 18H18.6a2.5 2.5 0 0 0 2.437-1.943l.913-4A2.5 2.5 0 0 0 19.513 9Zm-9.137 2.013a.5.5 0 0 1 .111-.013h9.026a.5.5 0 0 1 .487.611l-.913 4A.5.5 0 0 1 18.6 16h-7.2a.5.5 0 0 1-.487-.389l-.913-4a.5.5 0 0 1 .376-.598Z" clip-rule="evenodd"/><path d="M6.49 5H5a1 1 0 0 1 0-2h2.29a1 1 0 0 1 .975.78l2.71 12a1 1 0 1 1-1.95.44L6.49 5ZM13 20.25a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Zm7 0a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCartCircleFilled0)"/></g>`),
		g.Group(children),
	)
}