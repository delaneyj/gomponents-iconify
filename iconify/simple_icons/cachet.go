package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cachet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.746.254C5.265.254 0 5.519 0 12c0 6.481 5.265 11.746 11.746 11.746c6.482 0 11.746-5.265 11.746-11.746c0-1.44-.26-2.82-.734-4.097l-.264-.709l-1.118 1.118l.1.288c.373 1.064.575 2.207.575 3.4a10.297 10.297 0 0 1-10.305 10.305A10.297 10.297 0 0 1 1.441 12A10.297 10.297 0 0 1 11.746 1.695c1.817 0 3.52.47 5.002 1.293l.32.178l1.054-1.053l-.553-.316A11.699 11.699 0 0 0 11.746.254zM22.97.841L9.05 14.761L5.328 11.04l-1.031 1.03l4.752 4.753L24 1.872z"/>`),
		g.Group(children),
	)
}