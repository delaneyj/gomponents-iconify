package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeReload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.549 1.046c-3.859 0-6.819 3.192-7.166 6.985H1.059l1.892 1.952l2.065-1.952H3.677c.331-3.229 2.747-5.958 5.937-5.958c3.412 0 6.189 2.888 6.189 6.437c0 3.549-2.777 6.438-6.189 6.438c-1.695 0-3.232-.713-4.35-1.865l-.821.826a7.364 7.364 0 0 0 5.106 2.065c4.092 0 7.419-3.349 7.419-7.464s-3.327-7.464-7.419-7.464z"/><path d="M9 3.99V9h3.96V8H9.97V3.99H9z"/></g>`),
		g.Group(children),
	)
}