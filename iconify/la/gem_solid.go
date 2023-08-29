package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GemSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.531 6L4.22 12.375l-.5.594l.5.656l11 14l.781 1l.781-1l11-14l.5-.656l-.5-.594L22.47 6zm.938 2h3.656l-2.688 4H7.125zm7.406 0h3.656l3.344 4h-4.313zM16 8.844L18.125 12h-4.25zM7.031 14h4.219l2.375 8.406zm6.282 0h5.343L16 23.313zm7.437 0h4.219l-6.594 8.375z"/>`),
		g.Group(children),
	)
}