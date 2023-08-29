package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fontstencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M769.113 256q-12 0-27-16.5t-30-39t-43.5-52.5t-60.5-51q-11-8-21.5-29.5t-10.5-34.5q0-14 9.5-23t22.5-9q16 0 79.5-.5t81.5-.5q15 0 31 16t24 32l9 16v128q0 26-19 45t-45 19zm-129 704q0 27-18.5 45.5t-45.5 18.5h-320q-26 0-45-18.5t-19-45.5q0-11 20-25.5t44-27.5t44-33t20-42V64q0-36 25.5-50t71-14t70.5 14t25 50v768q0 22 20 42t44 33t44 27.5t20 25.5zm-416-864q-32 22-60 52t-43 52.5t-30 39t-27 16.5q-27 0-45.5-19t-18.5-45V64q3-7 9-17.5t22.5-28.5t32.5-18h160q14 0 23 9.5t9 22.5q0 14-10.5 35.5t-21.5 28.5z"/>`),
		g.Group(children),
	)
}