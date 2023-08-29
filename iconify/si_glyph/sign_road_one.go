package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignRoadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m13.447 3.902l3.365-1.34l-3.365-1.491H9.959V.167h-.902v.904H5.101v2.831h3.956v2.157H4.566L1.083 7.531l3.483 1.412h4.491v6.031h.902V8.943h3.017V6.059H9.959V3.902h3.488z"/>`),
		g.Group(children),
	)
}