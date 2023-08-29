package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShekelSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 5v22h2V7h1c3.879 0 7 3.121 7 7v7h2v-7c0-4.957-4.043-9-9-9zm15 0v20h-1c-3.879 0-7-3.121-7-7v-7h-2v7c0 4.957 4.043 9 9 9h3V5z"/>`),
		g.Group(children),
	)
}