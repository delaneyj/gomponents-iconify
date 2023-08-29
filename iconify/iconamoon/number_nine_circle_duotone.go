package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M10.164 16.452a1 1 0 0 0 1.672 1.096l-1.672-1.096Zm5.197-4.283a1 1 0 0 0-1.672-1.096l1.672 1.096ZM10 10a2 2 0 0 1 2-2V6a4 4 0 0 0-4 4h2Zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4v2Zm2 2a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4h-2Zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4v-2Zm-.164 5.548l3.525-5.38l-1.672-1.095l-3.525 5.379l1.672 1.096Z"/></g>`),
		g.Group(children),
	)
}