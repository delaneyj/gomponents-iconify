package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#295892" d="M30 17.832L2 23.28V30l28-5.447Zm0-10.579L2 12.7v6.72l28-5.447Z"/><path fill="#a1e0ff" d="M30 13.992L2 8.721V2l28 5.271Zm0 10.579L2 19.3v-6.72l28 5.27Z"/>`),
		g.Group(children),
	)
}