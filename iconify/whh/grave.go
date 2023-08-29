package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 1024H0q0-66 108.5-118T384 838V512H160q-13 0-22.5-9.5T128 480V288q0-13 9.5-22.5T160 256h224V32q0-13 9.5-22.5T416 0h192q13 0 22.5 9.5T640 32v224h224q13 0 22.5 9.5T896 288v192q0 13-9.5 22.5T864 512H640v326q167 16 275.5 68t108.5 118z"/>`),
		g.Group(children),
	)
}