package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBenin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#83bf4f" d="M22 60.3V3.7C10.4 7.8 2 18.9 2 32c0 13.1 8.4 24.2 20 28.3z"/><path fill="#ffce31" d="M22 32h40C62 15.4 48.6 2 32 2c-3.5 0-6.9.6-10 1.7V32"/><path fill="#ed4c5c" d="M22 32v28.3c3.1 1.1 6.5 1.7 10 1.7c16.6 0 30-13.4 30-30H22"/>`),
		g.Group(children),
	)
}