package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentRss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.958.042H2.02v15.911h11.959V4.042H7.958v-4zm-3.979 13.02v-2.305a2.305 2.305 0 0 1 2.306 2.305H3.979zm3.615 0c0-2.07-1.57-3.659-3.614-3.659V7.83c2.718 0 5.177 2.48 5.177 5.232H7.594zm4.485-.025h-1.562c0-3.686-2.835-6.583-6.507-6.583V4.893c4.367.001 8.069 3.76 8.069 8.144z"/><path d="M9.047.036v2.953h4.914L9.047.036z"/></g>`),
		g.Group(children),
	)
}