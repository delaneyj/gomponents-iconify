package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ftpsession(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.678 768h-320v128h128q26 0 45 19t19 45.5t-19 45t-45 18.5h-384q-26 0-45-18.5t-19-45t18.5-45.5t45.5-19h128V768h-320q-53 0-90.5-37.5T.678 640V256q0-26 19-45t45-19h480l46-85q7-18 28-30.5t44-12.5h240q24 0 45 12.5t28 30.5q49 58 49 85v448q0 53-37.5 90.5t-90.5 37.5zm-384-512q-87 0-160.5 43t-116.5 116.5t-43 160.5q0 33 6 64h66q-8-32-8-64q0-96 64-169v9l131 177q5 21 22 34t39 13q24 0 41.5-15t21.5-38l65-107v-32h-32l-93 57l-163-121h-9q73-64 169-64q106 0 181 75t75 181q0 31-8 64h66q6-32 6-64q0-87-43-160.5T673.178 299t-160.5-43zm-466-212q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84h-480z"/>`),
		g.Group(children),
	)
}