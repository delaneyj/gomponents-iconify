package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBurkinaFaso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#83bf4f" d="M32 62c16.6 0 30-13.4 30-30H2c0 16.6 13.4 30 30 30z"/><path fill="#ed4c5c" d="M32 2C15.4 2 2 15.4 2 32h60C62 15.4 48.6 2 32 2z"/><path fill="#ffce31" d="m32 36.8l5.6 4.2l-2.1-6.9l5.5-4.2h-6.9L32 23l-2.1 6.9H23l5.5 4.2l-2.1 6.9z"/>`),
		g.Group(children),
	)
}