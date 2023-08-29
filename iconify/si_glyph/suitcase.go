package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3v12.969h7.99V3H5zm9.037.031V16h.711C15.992 16 17 15.021 17 13.812V5.216c0-1.205-1.008-2.185-2.252-2.185h-.711zM1 5.217v8.596c0 1.208 1.018 2.188 2.276 2.188h.685V3.032h-.685C2.018 3.031 1 4.011 1 5.217zM9.968.047H8.019c-1.25 0-2.009.871-2.009 1.944h.961c0-.588.327-1.037 1.048-1.037h1.949c.719 0 1.057.45 1.057 1.037h.957c0-1.073-.763-1.944-2.014-1.944z"/>`),
		g.Group(children),
	)
}