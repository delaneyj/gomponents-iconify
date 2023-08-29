package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFrance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#428bc1" d="M1.9 32c0 13.1 8.4 24.2 20 28.3V3.7C10.3 7.8 1.9 18.9 1.9 32z"/><path fill="#ed4c5c" d="M61.9 32c0-13.1-8.3-24.2-20-28.3v56.6c11.7-4.1 20-15.2 20-28.3"/><path fill="#fff" d="M21.9 60.3c3.1 1.1 6.5 1.7 10 1.7s6.9-.6 10-1.7V3.7C38.8 2.6 35.5 2 31.9 2s-6.9.6-10 1.7v56.6"/>`),
		g.Group(children),
	)
}