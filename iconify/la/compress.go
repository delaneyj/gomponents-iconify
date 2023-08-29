package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 4v7H4v2h9V4h-2zm8 0v9h9v-2h-7V4h-2zM4 19v2h7v7h2v-9H4zm15 0v9h2v-7h7v-2h-9z"/>`),
		g.Group(children),
	)
}