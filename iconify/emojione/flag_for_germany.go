package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGermany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#3e4347" d="M31.9 2C18.8 2 7.7 10.4 3.6 22h56.6C56.1 10.4 45 2 31.9 2z"/><path fill="#ffe62e" d="M31.9 62c13.1 0 24.2-8.3 28.3-20H3.6c4.1 11.7 15.2 20 28.3 20z"/><path fill="#ed4c5c" d="M3.6 22c-1.1 3.1-1.7 6.5-1.7 10s.6 6.9 1.7 10h56.6c1.1-3.1 1.7-6.5 1.7-10s-.6-6.9-1.7-10H3.6"/>`),
		g.Group(children),
	)
}