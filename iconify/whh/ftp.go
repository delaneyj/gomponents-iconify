package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ftp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.678 768h-320v128h128q26 0 45 19t19 45.5t-19 45t-45 18.5h-384q-27 0-45.5-18.5t-18.5-45t19-45.5t45-19h128V768h-320q-53 0-90.5-37.5T.678 640V256q0-26 19-45t45-19h480l46-84q7-19 28-31.5t44-12.5h240q24 0 45 12.5t28 31.5l49 84v448q0 53-37.5 90.5t-90.5 37.5zm-454-294l-135-145q-8-9-19-9t-19 9l-135 145q-8 9-5 23.5t11 14.5h52v96q0 13 9.5 22.5t22.5 9.5h128q13 0 22.5-9.5t9.5-22.5v-96h52q8 0 11-14.5t-5-23.5zm442-26h-52v-96q0-13-9.5-22.5t-22.5-9.5h-128q-13 0-22.5 9.5t-9.5 22.5v96h-52q-8 0-11 14.5t5 23.5l135 145q8 9 19 9t19-9l135-145q8-9 5-23.5t-11-14.5zm-838-404q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84h-480z"/>`),
		g.Group(children),
	)
}