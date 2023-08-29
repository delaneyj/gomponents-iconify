package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentUrgent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1408H731l-475 475v-475H0V128h2048zm-128 128H128v1152h256v293l293-293h1243V256zm-896 640H896V384h128v512zm0 256H896v-128h128v128z"/>`),
		g.Group(children),
	)
}