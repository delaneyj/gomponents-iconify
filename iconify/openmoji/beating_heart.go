package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeatingHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFA7C0" d="M59.5 25c0-6.904-5.596-12.5-12.5-12.5a12.497 12.497 0 0 0-11 6.56a12.497 12.497 0 0 0-11-6.56c-6.904 0-12.5 5.596-12.5 12.5c0 2.97 1.04 5.694 2.77 7.839l-.004.003L36 58.54l20.734-25.698l-.004-.003A12.44 12.44 0 0 0 59.5 25z"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path stroke-linecap="round" d="M8 22s0-9 8-12M4 19s0-9 8-12m52 15s0-9-8-12m12 9s0-9-8-12"/><path d="M59.5 25c0-6.904-5.596-12.5-12.5-12.5a12.497 12.497 0 0 0-11 6.56a12.497 12.497 0 0 0-11-6.56c-6.904 0-12.5 5.596-12.5 12.5c0 2.97 1.04 5.694 2.77 7.839l-.004.003L36 58.54l20.734-25.698l-.004-.003A12.44 12.44 0 0 0 59.5 25z"/></g>`),
		g.Group(children),
	)
}