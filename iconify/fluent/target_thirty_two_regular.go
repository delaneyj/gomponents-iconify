package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-7-2a7 7 0 1 1 14 0a7 7 0 0 1-14 0Zm7-5a5 5 0 1 0 0 10a5 5 0 0 0 0-10ZM4 16.001C4 9.373 9.373 4 16.001 4c6.628 0 12.002 5.373 12.002 12.001c0 6.628-5.373 12.002-12.002 12.002C9.373 28.003 4 22.63 4 16ZM16.001 6C10.478 6 6 10.478 6 16.001c0 5.524 4.478 10.002 10.001 10.002c5.524 0 10.002-4.478 10.002-10.002C26.003 10.478 21.525 6 16 6Z"/>`),
		g.Group(children),
	)
}