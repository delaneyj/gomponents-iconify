package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nexus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M464 1024H304q-144 0-224-31.5T0 896V128q0-65 80-96.5T304 0h160q145 0 224.5 31.5T768 128v768q0 65-80 96.5T464 1024zm240-896H64v768h640V128zM128 832V191h384z"/>`),
		g.Group(children),
	)
}