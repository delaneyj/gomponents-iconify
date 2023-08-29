package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#A3BAF0" d="m125.306 214.373l.934 77.731l1.68 3.46l127.679-73.75l-72.306-41.811l-55.373 31.94l-2.614 2.43Z"/><path fill="#4474E1" d="m.24 221.813l127.68 73.75v-83.54l-55.293-31.94v-67L4.955 74.393H.24v147.42Z"/><path fill="#C7D6F7" d="m2.966 75.967l67.378 38.918l2.283 1.318l55.293-31.94l55.373 31.94l72.306-41.81L127.919.641L.242 74.392l2.725 1.575Z"/>`),
		g.Group(children),
	)
}