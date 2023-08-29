package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperClip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.346 16C5.009 16 4 14.907 4 13.725V3.998C4 1.634 5.254 0 7.698 0h.67C11.045 0 12 1.56 12 3.998v7.007h-.954V3.998C11.046 2.414 10.409 1 8.367 1h-.685C5.872 1 5 2.318 5 3.998v9.727c0 .738.448 1.274 1.345 1.274h1.338c.852 0 1.379-.488 1.379-1.274V5.756c0-.531-.081-.716-1.119-.765c-1.059.052-.943.271-.943.765v4.254H5.999V5.756c0-1.121.636-1.696 1.944-1.758C9.249 4.058 10 4.616 10 5.756v7.969C10 14.947 8.966 16 7.682 16H6.346z"/>`),
		g.Group(children),
	)
}