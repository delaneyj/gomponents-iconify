package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlighLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.043 1.938c0 .518.42.938.938.938h14.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938zm0 12c0 .518.42.938.938.938h14.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938zm0-6c0 .518.42.938.938.938h10.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938zm0 3c0 .518.42.938.938.938h8.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938zm0-6c0 .518.42.938.938.938h6.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938z"/>`),
		g.Group(children),
	)
}