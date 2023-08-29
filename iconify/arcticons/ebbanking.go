package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ebbanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.83 18.77h10.52v10.52H31.83Zm-26.48 0h10.52v10.52H5.35Zm26.48 13.14h10.52v10.52H31.83Zm-13.2 0h10.52v10.52H18.63Zm-13.28 0h10.52v10.52H5.35ZM31.83 5.54h10.52v10.52H31.83Zm-13.2 0h10.52v10.52H18.63Zm-13.28 0h10.52v10.52H5.35Z"/>`),
		g.Group(children),
	)
}