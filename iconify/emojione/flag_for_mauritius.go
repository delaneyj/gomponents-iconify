package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMauritius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M32 2C21.3 2 11.9 7.6 6.6 16h50.7C52.1 7.6 42.7 2 32 2z"/><path fill="#2a5f9e" d="M6.6 16C3.7 20.6 2 26.1 2 32h60c0-5.9-1.7-11.4-4.6-16H6.6z"/><path fill="#699635" d="M6.6 48c5.3 8.4 14.7 14 25.4 14s20.1-5.6 25.4-14H6.6z"/><path fill="#ffce31" d="M57.4 48c2.9-4.6 4.6-10.1 4.6-16H2c0 5.9 1.7 11.4 4.6 16h50.8z"/>`),
		g.Group(children),
	)
}