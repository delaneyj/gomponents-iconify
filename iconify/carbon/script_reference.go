package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScriptReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm23.17 6l-2.58 2.58L26 30l4-4l-4-4l-1.42 1.41L27.17 26zm-8.34 0l2.58-2.58L20 22l-4 4l4 4l1.42-1.41L18.83 26zM25.7 9.3l-7-7A.909.909 0 0 0 18 2H8a2.006 2.006 0 0 0-2 2v12h2V4h8v6a2.006 2.006 0 0 0 2 2h6v6h2v-8a.91.91 0 0 0-.3-.7zM18 10V4.4l5.6 5.6z"/>`),
		g.Group(children),
	)
}