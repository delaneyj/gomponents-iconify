package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPhotoCameraCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.172 8.232L8.762 9.5H7.5A2.5 2.5 0 0 0 5 12v6a2.5 2.5 0 0 0 2.5 2.5h11A2.5 2.5 0 0 0 21 18v-6a2.5 2.5 0 0 0-2.5-2.5h-1.263l-.409-1.268a2.5 2.5 0 0 0-2.38-1.732h-2.897a2.5 2.5 0 0 0-2.38 1.732ZM7.5 10.5h1.99l.633-1.96a1.5 1.5 0 0 1 1.428-1.04h2.898a1.5 1.5 0 0 1 1.427 1.04l.633 1.96H18.5A1.5 1.5 0 0 1 20 12v6a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 6 18v-6a1.5 1.5 0 0 1 1.5-1.5Z"/><path d="M10 14.5a3 3 0 1 0 6 0a3 3 0 0 0-6 0Zm5 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPhotoCameraCircleFilled0)"/></g>`),
		g.Group(children),
	)
}