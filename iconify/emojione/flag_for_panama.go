package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPanama(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M32 2v30h30C62 15.4 48.6 2 32 2"/><path fill="#2a5f9e" d="M32 62V32H2c0 16.6 13.4 30 30 30z"/><path fill="#f9f9f9" d="M32 62V32h30c0 16.6-13.4 30-30 30m0-60v30H2C2 15.4 15.4 2 32 2z"/><path fill="#2a5f9e" d="m17 20.7l3.1 2.3l-1.2-3.8l3.1-2.4h-3.8L17 13l-1.2 3.8H12l3.1 2.4l-1.2 3.8z"/><path fill="#ed4c5c" d="m47 46.7l3.1 2.3l-1.2-3.8l3.1-2.4h-3.8L47 39l-1.2 3.8H42l3.1 2.4l-1.2 3.8z"/>`),
		g.Group(children),
	)
}