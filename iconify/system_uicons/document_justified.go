package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentJustified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<defs><path id="systemUiconsDocumentJustified0" d="M16.5 15.5v-10a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm-9-8h6m-6 3h6m-6 3h6"/></defs><g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><use href="#systemUiconsDocumentJustified0"/><use href="#systemUiconsDocumentJustified0"/></g>`),
		g.Group(children),
	)
}