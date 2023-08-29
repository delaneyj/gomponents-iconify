package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageRuby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.8 2.07c2.52.43 3.24 2.16 3.2 3.97V6l-1.14 14.93l-14.78 1.01h.01c-1.23-.05-3.96-.17-4.09-3.99l1.37-2.5l2.77 6.46l2.36-7.62l-.05.01l.02-.02l7.71 2.46l-1.99-7.78l7.35-.46l-5.79-4.74l3.05-1.7v.01M2 17.91v.02v-.02M6.28 6.23c2.96-2.95 6.79-4.69 8.26-3.2c1.46 1.47-.08 5.09-3.04 8.03c-3 2.94-6.77 4.78-8.24 3.3c-1.47-1.49.04-5.19 3.01-8.13h.01Z"/>`),
		g.Group(children),
	)
}