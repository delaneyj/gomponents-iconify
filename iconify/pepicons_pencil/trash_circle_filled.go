package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTrashCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M11.5 17.999a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2 0a.5.5 0 1 1-1 0v-6a.5.5 0 0 1 1 0v6Zm-1-10.5h-3a1.501 1.501 0 0 1 3-.001Z"/><path d="M7.5 7.999a.5.5 0 1 1 0-1h11a.5.5 0 0 1 0 1h-11Z"/><path fill-rule="evenodd" d="M17.5 8.5h-9A.5.5 0 0 0 8 9v11a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V9a.5.5 0 0 0-.5-.5ZM9 19.5v-10h8v10H9Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTrashCircleFilled0)"/></g>`),
		g.Group(children),
	)
}