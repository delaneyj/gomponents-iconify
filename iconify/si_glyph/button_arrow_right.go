package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.969 8.041c0 4.389-3.582 7.959-7.986 7.959c-4.4 0-7.982-3.57-7.982-7.959c0-4.389 3.582-7.959 7.982-7.959c4.404 0 7.986 3.57 7.986 7.959zm-14 0c0 3.305 2.712 5.996 6.045 5.996c3.333 0 6.047-2.691 6.047-5.996c0-3.305-2.713-5.996-6.047-5.996c-3.334 0-6.045 2.691-6.045 5.996z"/><path d="m9.057 10.969l3.904-3.002l-3.905-2.951v2.003H5.793c-.346 0-.791.324-.791.955c0 .63.483.975.826.975h3.229v2.02z"/></g>`),
		g.Group(children),
	)
}