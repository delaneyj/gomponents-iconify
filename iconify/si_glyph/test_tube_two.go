package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.726.04L8.32 1.443l.703.703l-8.48 8.323c-.738.735-.674 1.996.143 2.811l2.031 2.025c.816.814 2.047.91 2.785.173l8.514-8.353l.684.684l1.408-1.404L9.726.04zm-6.37 14.358l-1.771-1.771c-.502-.501-.591-1.232-.195-1.623l8.35-8.225l3.59 3.592l-8.351 8.222c-.395.395-1.123.305-1.623-.195z"/><path d="M5.273 12.78c-.34.336-.904.316-1.406-.184l-.376-.377c-.503-.504-.636-1.184-.298-1.518l3.858-3.858l3.203.909l-4.981 5.028z"/></g>`),
		g.Group(children),
	)
}