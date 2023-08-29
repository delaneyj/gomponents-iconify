package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M272 227.616V121.388h-32v106.229l-98.623-59.682l-2.431 4.017v31.915L225.096 256l-86.15 52.133v31.915l2.431 4.017L240 284.383v106.229h32V284.384l98.623 59.681l2.431-4.016v-31.915L286.903 256l86.151-52.134v-31.915l-2.431-4.016L272 227.616z"/><path fill="currentColor" d="M425.857 87.379A239.365 239.365 0 0 0 87.344 425.892A239.365 239.365 0 0 0 425.857 87.379ZM256.6 464c-114.341 0-207.364-93.023-207.364-207.364S142.259 49.271 256.6 49.271s207.365 93.023 207.365 207.365S370.942 464 256.6 464Z"/>`),
		g.Group(children),
	)
}