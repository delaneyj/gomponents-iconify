package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578V6.423L12 .845Zm0 2.31L5.34 7L12 10.845L18.66 7L12 3.155Zm7.66 5.577L13 12.577v7.69l6.66-3.844V8.732ZM11 20.268v-7.69L4.34 8.731v7.69L11 20.269ZM13 5v4h-2V5h2Zm-3.098 9.366l-3.464 2l-1-1.732l3.464-2l1 1.732Zm5.196-1.732l3.464 2l-1 1.732l-3.464-2l1-1.732Z"/>`),
		g.Group(children),
	)
}