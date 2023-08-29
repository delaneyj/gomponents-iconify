package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTaiwan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M32 2v30H2c0 16.6 13.4 30 30 30s30-13.4 30-30S48.6 2 32 2z"/><path fill="#2a5f9e" d="M32 2C15.4 2 2 15.4 2 32h30V2z"/><path fill="#fff" d="m24 20.3l5-1.3l-5-1.3l3.7-3.7l-5 1.3l1.3-5l-3.7 3.7L19 9l-1.3 5l-3.7-3.7l1.3 5l-5-1.3l3.7 3.7L9 19l5 1.3l-3.7 3.7l5-1.3l-1.3 5l3.7-3.7l1.3 5l1.3-5l3.7 3.7l-1.3-5l5 1.3l-3.7-3.7"/><circle cx="19" cy="19" r="5.7" fill="#428bc1"/><circle cx="19" cy="19" r="5" fill="#fff"/>`),
		g.Group(children),
	)
}