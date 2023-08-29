package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c8.837 0 16 7.163 16 16s-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0zm-.343 5.108l-.084.07l-10.394 10.39a.607.607 0 0 0-.071.775l.07.085l10.39 10.393a.607.607 0 0 0 .776.071l.084-.07l10.39-10.39a.608.608 0 0 0 .004-.864l-10.39-10.39a.607.607 0 0 0-.775-.07zm.387 2.328l.04.027l8.454 8.453a.117.117 0 0 1 .027.131l-.027.041l-7.616 7.616V16h-1.848v7.704l-7.616-7.616a.12.12 0 0 1-.027-.131l.027-.04l8.454-8.454a.12.12 0 0 1 .132-.027z"/>`),
		g.Group(children),
	)
}