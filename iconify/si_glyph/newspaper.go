package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.031 2.046v9.788c0 .386-1.062.389-1.062.016V6.063H1.022v6.918c-.004.104.031.991.979.991l14-.012s.973-.049.973-.96V2l-12.943.046zm4.011 8.016h-2.07V6.937h2.07v3.125zM15 8H9V6.958h6V8zm0 2.083H8.981V8.939l6.019.02v1.124zm0 1.955H5.987v-1.08H15v1.08zm0-5.997H6V3.937h9v2.104z"/>`),
		g.Group(children),
	)
}