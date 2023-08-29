package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1004.19 1003q-19 19-45 20t-44-17l-95-95q-86 57-191.5 46t-180.5-86l-16-16l424-423l15 15q75 75 86 180.5t-45 192.5l95 94q18 18 17 44t-20 45zm-473-573l-102-101l-100 100l102 101q18 18 17 44t-20 45t-45 20t-44-17l-102-101l-70 70l-15-16q-75-75-86-180.5t45-191.5l-94-94q-18-18-17-44t20-45t45-20t44 17l94 94q86-57 191.5-46t181.5 86l15 16l-70 70l102 101q18 18 17 44t-20 45t-45 20t-44-17z"/>`),
		g.Group(children),
	)
}