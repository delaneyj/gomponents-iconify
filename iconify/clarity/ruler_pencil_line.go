package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerPencilLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M9 17.41V27h9.59l-2-2H11v-5.59l-2-2z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M34.87 32.29L32 29.38V32H4v-4.15h2v-1.6H4V19.6h2V18H4v-6.4h2V10H4V4.41l15.94 15.85v-2.82L3.71 1.29A1 1 0 0 0 2 2v31a1 1 0 0 0 1 1h31.16a1 1 0 0 0 .71-1.71Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M24 30h4a2 2 0 0 0 2-2V8.7l-2.3-4.23a2 2 0 0 0-1.76-1a2 2 0 0 0-1.76 1.08L22 8.72V28a2 2 0 0 0 2 2Zm0-20.8l1.94-3.77L28 9.21V24h-4Zm0 16.43h4v2.44h-4Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}