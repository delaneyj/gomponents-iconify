package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkypeClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1431 1705q28 28 43 65t15 77q0 42-15 78t-43 64t-63 43t-79 16q-83 0-142-59l-594-593q-28-28-43-65t-15-77V201q0-42 15-78t43-64t63-43t79-16q42 0 78 16t64 43t43 63t16 79v970l535 534z"/>`),
		g.Group(children),
	)
}