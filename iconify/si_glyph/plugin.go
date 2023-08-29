package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plugin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.016 7.038h-5v2.931h5v.829l5.922 2.175v-8.93L5.016 6.219v.819zm10.937-.054v-.968h-2.971V4.043h-.962v8.93h.962v-1.989h2.971v-.968h-2.971V6.984h2.971z"/>`),
		g.Group(children),
	)
}