package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowCascade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 14.6c-.967 0-1.796.576-2.176 1.4H8.5A1.5 1.5 0 0 1 7 14.5v-3.85c.456.218.961.35 1.5.35h3.324a2.396 2.396 0 0 0 4.576-1a2.397 2.397 0 0 0-4.576-1H8.5A1.5 1.5 0 0 1 7 7.5V5.176A2.396 2.396 0 0 0 6 .6a2.397 2.397 0 0 0-1 4.576V14.5A3.5 3.5 0 0 0 8.5 18h3.324a2.396 2.396 0 0 0 4.576-1a2.4 2.4 0 0 0-2.4-2.4zm0-5.985a1.384 1.384 0 1 1 0 2.768a1.384 1.384 0 0 1 0-2.768zm-8-7a1.384 1.384 0 1 1 0 2.768a1.384 1.384 0 0 1 0-2.768zm8 16.77a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77z"/>`),
		g.Group(children),
	)
}