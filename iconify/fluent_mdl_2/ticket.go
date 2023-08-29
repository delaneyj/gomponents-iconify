package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 896q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10v512H0v-512q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10V384h2048v512zm-128-384H128v290q60 35 94 94t34 128q0 69-34 128t-94 94v290h1792v-290q-60-35-94-94t-34-128q0-69 34-128t94-94V512z"/>`),
		g.Group(children),
	)
}