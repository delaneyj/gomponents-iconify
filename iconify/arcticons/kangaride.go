package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kangaride(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.737 19.448l-8.105 24l14.789.052l.263-10.579l8.421 1.263l3.263-3.158l-18.63-11.578Zm.947-2.159c4.876-3.15 6.607-6.525 4.58-12.789l-4.58 12.79Zm14.474-7.21c-.576 6.893-8.738 8.947-11.58 8.053l11.58-8.053Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.875 27.783a1.838 1.838 0 1 1 .005-.001"/>`),
		g.Group(children),
	)
}