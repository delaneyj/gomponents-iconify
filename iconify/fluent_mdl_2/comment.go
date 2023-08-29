package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1408H731l-475 475v-475H0V128h2048zm-128 128H128v1152h256v293l293-293h1243V256zm-640 640h-128V640h256v256q0 27-10 50t-27 40t-41 28t-50 10V896zm-512 0H640V640h256v256q0 27-10 50t-27 40t-41 28t-50 10V896z"/>`),
		g.Group(children),
	)
}