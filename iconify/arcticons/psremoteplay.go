package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psremoteplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.56 21.993h0a7.274 7.274 0 0 0-12.23-4.337H18.66a7.284 7.284 0 0 0-4.997-1.998h0a7.294 7.294 0 0 0-7.244 6.335h0L4.511 37.22c-.32 3.886 6.524 3.996 7.803 1.319l3.088-4.897a4.316 4.316 0 0 1 3.587-2.298h9.992a4.237 4.237 0 0 1 3.617 2.298l3.087 4.896c1.28 2.658 8.124 2.568 7.804-1.318Z"/><circle cx="23.985" cy="12.853" r=".75" fill="currentColor"/><circle cx="23.985" cy="8.379" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}