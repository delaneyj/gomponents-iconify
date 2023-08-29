package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShakeDevice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1897 1527q0 34-18 63t-49 45l-804 400q-13 6-26 9t-27 4q-35 0-63-17t-45-50L164 574q-13-28-13-53q0-34 18-63t49-45l804-400q28-13 53-13q35 0 64 18t44 49l701 1407q13 28 13 53zm-921 390l790-393l-694-1393l-790 393l694 1393zm427-383l-201 100l-57-115l201-100l57 115zm261-1187l-83 82l-90-90l237-237l237 237l-90 90l-83-82v293h-128V347zM384 1701l83-82l90 90l-237 237l-237-237l90-90l83 82v-293h128v293z"/>`),
		g.Group(children),
	)
}