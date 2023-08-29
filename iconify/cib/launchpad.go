package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Launchpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.333 24.901l-5.912-3.407l-6.375 3.683V32zM14.953 32v-6.823l-6.38-3.688l-5.907 3.412zm1.042-8.636l6.385-3.681v-7.371L16 8.636l-6.38 3.676v7.371zM1.625 8.912v14.187l5.911-3.411v-7.371L1.625 8.906zm26.156 1.495l-3.312 1.911v7.365l5.905 3.411V8.907zM2.667 7.099l5.917 3.412l6.375-3.683V.005zm20.754 3.412l5.912-3.412L17.041 0v6.828z"/>`),
		g.Group(children),
	)
}