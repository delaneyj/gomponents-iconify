package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 6.792c-.39-.193-.847-.089-1.023.231c0 0-1.096 2.399-1.734 2.297c-.344-.056-.619-.38-.742-1.005V1.771c0-.426-.439-.771-.98-.771c-.54 0-1.02.346-1.02.771v5.167h-1V.771c0-.426-.439-.771-.98-.771c-.54 0-1.02.346-1.02.771v6.167h-1V2.771c0-.426-.44-.771-.981-.771c-.54 0-1.019.346-1.019.771v8.415c0 2.584 1.729 4.721 5.678 4.721c4.883 0 6.205-8.188 6.205-8.188c.174-.32.003-.736-.384-.927z"/>`),
		g.Group(children),
	)
}