package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.148 5.25H5.335l-1.18-2.115A.75.75 0 0 0 3.5 2.75H2a.75.75 0 0 0 0 1.5h1.06l1.164 2.088L6.91 12.28l.003.006l.237.523l-2.697 2.877a.75.75 0 0 0 .462 1.258l2.458.281a40.68 40.68 0 0 0 9.254 0l2.458-.28a.75.75 0 0 0-.17-1.491l-2.458.28a39.256 39.256 0 0 1-8.914 0l-.975-.11l1.98-2.112a.768.768 0 0 0 .053-.064l.752.098c1.055.138 2.122.165 3.182.08a9.29 9.29 0 0 0 6.366-3.268l.579-.685a.734.734 0 0 0 .053-.072l1.078-1.642c.764-1.164-.071-2.71-1.463-2.71ZM6.5 18.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM16 20a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		g.Group(children),
	)
}