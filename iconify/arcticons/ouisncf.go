package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ouisncf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.25 35.37H32.6a14.25 14.25 0 1 0-17.2 0H5.76a21.5 21.5 0 1 1 36.49 0Z"/>`),
		g.Group(children),
	)
}