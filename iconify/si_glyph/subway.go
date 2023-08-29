package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M5 15.979h-.979l1-4H6l-1 4zm8 0h-.979l-1-4H12l1 4z"/><path d="M5 14h6.9v.926H5zM11.988.038H5.095c-2.242 0-3.054.688-3.054 2.775v7.784c0 1.848.813 2.441 2.984 2.441h7.725c1.574 0 2.291-.602 2.291-2.44V2.812c0-2.087-.81-2.775-3.053-2.775zM6.213 1h4.574c.117 0 .213.226.213.5c0 .273-.096.5-.213.5H6.213C6.096 2 6 1.773 6 1.5c0-.274.096-.5.213-.5zm-.197 10.016H3.969V8.985h2.047v2.031zm7 0h-2.047V8.985h2.047v2.031zm.037-5.235c0 1.064-.166 1.257-1.128 1.257H5.091c-.964 0-1.127-.146-1.127-1.21V4.156c0-1.062.163-1.198 1.127-1.198h6.834c.962 0 1.128.198 1.128 1.261v1.562z"/></g>`),
		g.Group(children),
	)
}