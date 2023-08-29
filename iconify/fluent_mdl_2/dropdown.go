package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h2048v1408H0V256zm1920 1280V384H128v1152h1792zm-128-640l-256 256l-256-256h512z"/>`),
		g.Group(children),
	)
}