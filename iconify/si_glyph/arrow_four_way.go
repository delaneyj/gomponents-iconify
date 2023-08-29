package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFourWay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m16.844 7.747l-2.062-2.591a.551.551 0 0 0-.779 0v1.875H9.989V2.969h1.855a.558.558 0 0 0 0-.78L9.391.188a.541.541 0 0 0-.771 0L6.188 2.189a.554.554 0 0 0 0 .78H8.01v4.062H3.985V5.172a.551.551 0 0 0-.779 0S1.012 7.576 1.012 8.07c0 .428 2.194 2.68 2.194 2.68a.55.55 0 0 0 .779 0V8.984H8.01v4.047H6.172a.54.54 0 0 0 0 .76s2.332 2.192 2.826 2.192c.43 0 2.846-2.192 2.846-2.192a.54.54 0 0 0 0-.76H9.989V8.984h4.014v1.812a.551.551 0 0 0 .779 0l2.062-2.27a.551.551 0 0 0 0-.779z"/>`),
		g.Group(children),
	)
}