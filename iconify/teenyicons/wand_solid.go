package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WandSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 0v2H9V0h1ZM5.5.792L7.207 2.5l-.707.707L4.793 1.5L5.5.792Zm8.707.708L12.5 3.206l-.707-.707L13.5.792l.707.708ZM4 4.997h2v1H4v-1Zm9 0h2v1h-2v-1ZM7.207 8.495l-6.354 6.35l-.706-.708L6.5 7.787l.707.708Zm5.293-.707l1.707 1.706l-.707.707l-1.707-1.706l.707-.707ZM10 8.994v2H9v-2h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}