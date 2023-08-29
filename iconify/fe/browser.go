package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBrowser0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBrowser1" fill="currentColor" fill-rule="nonzero"><path id="feBrowser2" d="M4 10v8h16v-8H4Zm0-4v2h2V6H4Zm4 0v2h12V6H8ZM4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/></g></g>`),
		g.Group(children),
	)
}