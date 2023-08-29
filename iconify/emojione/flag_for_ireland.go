package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForIreland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#83bf4f" d="M1.8 32c0 13.1 8.4 24.2 20 28.3V3.7C10.1 7.8 1.8 18.9 1.8 32z"/><path fill="#ff8736" d="M61.8 32c0-13.1-8.4-24.2-20-28.3v56.6c11.6-4.1 20-15.2 20-28.3"/><path fill="#fff" d="M21.8 60.3c3.1 1.1 6.5 1.7 10 1.7s6.9-.6 10-1.7V3.7C38.6 2.6 35.3 2 31.8 2s-6.9.6-10 1.7v56.6"/>`),
		g.Group(children),
	)
}