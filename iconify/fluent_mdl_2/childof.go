package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Childof(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 256v128H128V256h896zm557 723l365 365l-365 365l-90-90l210-211H512V512h128v768h1061l-210-211l90-90z"/>`),
		g.Group(children),
	)
}