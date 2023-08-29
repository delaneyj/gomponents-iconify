package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForStLucia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#49c3f2"/><path fill="#fff" d="M32 15L17 45h30z"/><path fill="#3e4347" d="M32 21L20 45h24z"/><path fill="#ffce31" d="M32 33L20 45h24z"/>`),
		g.Group(children),
	)
}