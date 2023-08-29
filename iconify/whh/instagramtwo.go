package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagramtwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-640-960h-128v256h128V64zm256 320q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136t-56-136t-136-56zm384-224q0-13-9.5-22.5t-22.5-9.5h-64q-13 0-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5v-64zm64 224h-192q64 86 64 192q0 87-43 160.5T672.928 853t-160.5 43t-160.5-43t-116.5-116.5t-43-160.5q0-106 64-192h-192v512q0 27 18.5 45.5t45.5 18.5h768q26 0 45-18.5t19-45.5V384z"/>`),
		g.Group(children),
	)
}