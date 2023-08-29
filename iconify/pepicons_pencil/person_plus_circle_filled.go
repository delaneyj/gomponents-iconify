package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPlusCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPersonPlusCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.95 9.75a.5.5 0 0 1 .5-.5h5.25a.5.5 0 1 1 0 1h-5.25a.5.5 0 0 1-.5-.5Z"/><path d="M18 13a.5.5 0 0 1-.5-.5V7.25a.5.5 0 0 1 1 0v5.25a.5.5 0 0 1-.5.5ZM9.8 6a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM6.3 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z"/><path d="M3.8 17.5c0-3.322 2.67-6.5 6-6.5s6 3.178 6 6.5v2a.5.5 0 0 1-1 0v-2c0-2.873-2.32-5.5-5-5.5s-5 2.627-5 5.5v2a.5.5 0 0 1-1 0v-2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPersonPlusCircleFilled0)"/></g>`),
		g.Group(children),
	)
}