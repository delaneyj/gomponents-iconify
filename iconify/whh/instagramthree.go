package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagramthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-768-960q-26 0-45 18.5t-19 45.5v192h64V64zm384 832q87 0 160.5-43t116.5-116.5t43-160.5t-43-160.5t-116.5-116.5t-160.5-43t-160.5 43t-116.5 116.5t-43 160.5t43 160.5t116.5 116.5t160.5 43zm448-768q0-27-18.5-45.5t-45.5-18.5h-640v226q110-98 256-98q84 0 158 34t128 94h162V128zm-192 96v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5zm-256 160q80 0 136 56t56 136t-56 136t-136 56t-136-56t-56-136t56-136t136-56z"/>`),
		g.Group(children),
	)
}