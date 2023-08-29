package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.98 7.984c0-3.763-3.16-6.659-7.012-6.953V.468c0-.258-.17-.467-.447-.467a.475.475 0 0 0-.49.467v.555c-3.941.205-7.018 3.139-7.018 6.962h.955V6.969H4.03v1.016h1.938V6.969h1.08v1.016h1.904V6.969h1.094l-.016 1.016h1.949V6.969h1.02v1.016h1.938V6.969h1.078v1.016h.965v-.001zm-7.96 6.218c0 .479-.507.816-1.003.816c-.316 0-1.035-.016-1.035-.494a.474.474 0 0 0-.48-.466a.473.473 0 0 0-.478.466c0 .99.97 1.473 1.993 1.473c1.025 0 1.978-.803 1.978-1.795v-6.17h-.974v6.17z"/>`),
		g.Group(children),
	)
}