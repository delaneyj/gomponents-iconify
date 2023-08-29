package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(.995 2.98)"><circle cx="7.958" cy="6.958" r="2.958"/><path d="M14.666 2.042h-3.713L9.937.031H6L5 2.042H1.333C.597 2.042 0 2.639 0 3.375v7.25c0 .736.597 1.334 1.333 1.334h13.333c.736 0 1.334-.598 1.334-1.334v-7.25c0-.736-.598-1.333-1.334-1.333zM6.953.969h2.094v1.062H6.953V.969zm1.049 10.064a4.052 4.052 0 0 1-4.055-4.048a4.052 4.052 0 0 1 4.055-4.048a4.051 4.051 0 0 1 4.055 4.048a4.051 4.051 0 0 1-4.055 4.048zM14 4.031h-2.047V2.969H14v1.062z"/></g>`),
		g.Group(children),
	)
}