package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManDoctor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.957 4.121c0-2.457-1.701-4.059-3.492-4.059c-1.794 0-3.457 1.586-3.457 4.059h-.001C5.014 5.779 6.569 8.92 8.483 8.92c1.913 0 3.464-3.141 3.474-4.799zm1.542 6.92a2.463 2.463 0 0 0-2.464 2.46a2.463 2.463 0 0 0 4.926 0a2.462 2.462 0 0 0-2.462-2.46zm.536 2.981V15h-1.053v-.978h-1.028v-1.033h1.028v-1.026h1.053v1.026h.989v1.033h-.989zm-4.637-1.02c0-1.241.545-2.354 1.406-3.109c.461-.406.434-.37.032-.175c-1.158.562-2.726.726-2.726.726s-2.908-.287-3.765-1.439C.24 9.005.116 14.969.116 14.969h9.775a4.16 4.16 0 0 1-.493-1.967z"/>`),
		g.Group(children),
	)
}