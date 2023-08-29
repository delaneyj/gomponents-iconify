package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><ellipse cx="4.486" cy="1.525" rx="1.486" ry="1.525"/><path d="m4.479 13.325l1.373-8.622s.029-.699-.636-.699c-1.501 0-5.29.004-5.29.004s3.979.713 3.979 2.059c0 0-1.308 5.76-1.308 7.29c0 1.531.836 2.644 2.337 2.644c1.225 0 2.338-1.336 2.254-3.283c-1.197 1.836-2.709 2.102-2.709.607z"/></g>`),
		g.Group(children),
	)
}