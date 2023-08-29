package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffce31" d="M8 50h48c2.6-3.5 4.5-7.6 5.4-12H2.6c.9 4.4 2.8 8.5 5.4 12"/><path fill="#83bf4f" d="M56 50H8c5.5 7.3 14.2 12 24 12s18.5-4.7 24-12"/><path fill="#ed4c5c" d="M2 32c0 2.1.2 4.1.6 6H32V2C15.4 2 2 15.4 2 32z"/><path fill="#83bf4f" d="M32 2v12h24C50.5 6.7 41.8 2 32 2z"/><path fill="#ffce31" d="M56 14H32v12h29.4c-.9-4.4-2.8-8.5-5.4-12"/><path fill="#83bf4f" d="M61.4 26H32v12h29.4c.4-1.9.6-3.9.6-6s-.2-4.1-.6-6"/><path fill="#fff" d="m18 27.3l6.8 4.7l-2.6-7.6l6.8-4.8h-8.4L18 12l-2.6 7.6H7l6.8 4.8l-2.6 7.6z"/>`),
		g.Group(children),
	)
}