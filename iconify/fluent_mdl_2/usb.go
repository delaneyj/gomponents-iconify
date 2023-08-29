package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func USB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 256h128v128H768V256zm384 0v128h-128V256h128zm256 256q27 0 50 10t40 27t28 41t10 50v1280q0 27-10 50t-27 40t-41 28t-50 10H512q-27 0-50-10t-40-27t-28-41t-10-50V640q0-27 10-50t27-40t41-28t50-10V0h896v512zm-768 0h640V128H640v384zm768 128H512v1280h896V640z"/>`),
		g.Group(children),
	)
}