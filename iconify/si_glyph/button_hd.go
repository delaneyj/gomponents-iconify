package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonHd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.276 2L9.015 3.062L3.723 2C2.218 2 .999 3.243.999 4.777v5.41c0 1.534 1.219 2.777 2.724 2.777l5.292-1.012l5.261 1.012c1.504 0 2.724-1.243 2.724-2.777v-5.41C17 3.243 15.78 2 14.276 2zm-6.231 8.032H6.957v-1.97H5.061v1.97H3.973V4.969h1.088v1.969h1.896V4.969h1.088v5.063zm3.955.029H9.953V4.96L12 4.959c2.303 0 2.072.812 2.072 2.552c0 1.738.335 2.55-2.072 2.55z"/><path d="M12.197 6.006H11v2.979h1.197c.99-.016.792-.248.792-1.488c0-1.243.199-1.477-.792-1.491z"/></g>`),
		g.Group(children),
	)
}