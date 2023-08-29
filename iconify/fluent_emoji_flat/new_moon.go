package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#533566" d="M30 16c0 7.732-6.268 14-14 14S2 23.732 2 16S8.268 2 16 2s14 6.268 14 14Z"/><path fill="#321B41" d="M15.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-6 9a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7ZM25 11.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM16.5 24a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-6.311 4.74a14.054 14.054 0 0 1-5.439-4.405a3.5 3.5 0 1 1 5.439 4.406Zm14.693-1.917a14.04 14.04 0 0 0 4.073-5.507a3.5 3.5 0 0 0-4.074 5.507Z"/></g>`),
		g.Group(children),
	)
}