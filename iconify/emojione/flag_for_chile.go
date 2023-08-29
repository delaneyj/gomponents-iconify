package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForChile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M2 32c0 16.6 13.4 30 30 30s30-13.4 30-30H2z"/><path fill="#f9f9f9" d="M32 2c16.6 0 30 13.4 30 30H32V2z"/><path fill="#2a5f9e" d="M32 2C15.4 2 2 15.4 2 32h30V2z"/><path fill="#f9f9f9" d="m20 21.7l4.9 3.3l-1.8-5.3l4.9-3.5h-6.1L20 11l-1.8 5.2H12l4.9 3.5l-1.8 5.3z"/>`),
		g.Group(children),
	)
}