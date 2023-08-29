package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milkshake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M304 149H197V98l96-57q17-11 9-30q-3-8-12.5-10T272 2L155 73v76H48q-17 0-30 13T5 192v64h26l38 218q1 16 14 27t29 11h132q16 0 28-10t15-26l34-220h26v-64q0-17-13-30t-30-13zM91 213v-21h42v21H91zm64-21h42v21h-42v-21zm64 0h42v21h-42v-21zm85 21h-21v-21h21v21zM48 192h21v21H48v-21zm196 277l-132-2l-38-211h204z"/>`),
		g.Group(children),
	)
}