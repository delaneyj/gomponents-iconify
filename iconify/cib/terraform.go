package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terraform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.041 5.688l9.912 5.031v10.073l-9.912-5.037zm11 5.031v10.073l9.917-5.037V5.688zM.047.068v10.068l9.912 5.036V5.104zm10.994 26.853l9.912 5.037V21.895l-9.912-5.036z"/>`),
		g.Group(children),
	)
}