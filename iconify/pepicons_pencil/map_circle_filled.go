package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilMapCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 11.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.5 9.286c0 2.673 2.653 7.214 4.5 7.214c1.848 0 4.5-4.541 4.5-7.214C17.5 6.65 15.493 4.5 13 4.5S8.5 6.65 8.5 9.286Zm8 0c0 2.193-2.348 6.214-3.5 6.214c-1.151 0-3.5-4.02-3.5-6.214C9.5 7.187 11.075 5.5 13 5.5s3.5 1.687 3.5 3.786Z" clip-rule="evenodd"/><path d="M16.435 12.14a.5.5 0 0 1 .369-.929a3 3 0 0 1 1.74 1.84l1.334 4A3 3 0 0 1 17.03 21H8.97a3 3 0 0 1-2.846-3.949l1.333-4a3 3 0 0 1 1.783-1.857a.5.5 0 1 1 .355.935a2 2 0 0 0-1.19 1.239l-1.333 4A2 2 0 0 0 8.97 20h8.062a2 2 0 0 0 1.897-2.633l-1.332-4a2 2 0 0 0-1.16-1.226Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilMapCircleFilled0)"/></g>`),
		g.Group(children),
	)
}