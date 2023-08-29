package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTonga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M26 2.6C12.3 5.4 2 17.5 2 32h24V2.6z"/><g fill="#c94747"><path d="M32 2c-2.1 0-4.1.2-6 .6V32H2c0 16.6 13.4 30 30 30s30-13.4 30-30S48.6 2 32 2z"/><path d="M22 18h-5v-5h-4v5H8v4h5v5h4v-5h5z"/></g>`),
		g.Group(children),
	)
}