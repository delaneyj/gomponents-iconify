package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.008 16.929c-4.385 0-7.95-3.563-7.95-7.942c0-4.381 3.565-7.944 7.95-7.944c4.384 0 7.95 3.563 7.95 7.944c0 4.378-3.566 7.942-7.95 7.942zM8.973 2.863c-3.354 0-6.084 2.742-6.084 6.111c0 3.369 2.73 6.111 6.084 6.111c3.355 0 6.085-2.742 6.085-6.111c0-3.369-2.73-6.111-6.085-6.111z"/><path d="M8 5h.959v3.978H8zm0 4h3.938v.979H8z"/></g>`),
		g.Group(children),
	)
}