package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGuinea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#75a843" d="M62 32c0 13.1-8.3 24.2-20 28.3V3.7C53.7 7.8 62 18.9 62 32z"/><path fill="#ed4c5c" d="M2 32C2 18.9 10.4 7.8 22 3.7v56.6C10.4 56.2 2 45.1 2 32z"/><path fill="#ffce31" d="M42 60.3c-3.1 1.1-6.5 1.7-10 1.7s-6.9-.6-10-1.7V3.7C25.1 2.6 28.5 2 32 2s6.9.6 10 1.7v56.6"/>`),
		g.Group(children),
	)
}