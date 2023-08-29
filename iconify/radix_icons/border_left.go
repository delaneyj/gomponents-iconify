package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.75 1v13H.25V1h1.5Z" clip-rule="evenodd"/><rect width="1" height="1" x="10" y="7" rx=".5" transform="rotate(90 10 7)"/><rect width="1" height="1" x="10" y="13" rx=".5" transform="rotate(90 10 13)"/><rect width="1" height="1" x="12" y="7" rx=".5" transform="rotate(90 12 7)"/><rect width="1" height="1" x="12" y="13" rx=".5" transform="rotate(90 12 13)"/><rect width="1" height="1" x="8" y="7" rx=".5" transform="rotate(90 8 7)"/><rect width="1" height="1" x="14" y="7" rx=".5" transform="rotate(90 14 7)"/><rect width="1" height="1" x="8" y="13" rx=".5" transform="rotate(90 8 13)"/><rect width="1" height="1" x="14" y="13" rx=".5" transform="rotate(90 14 13)"/><rect width="1" height="1" x="8" y="5" rx=".5" transform="rotate(90 8 5)"/><rect width="1" height="1" x="14" y="5" rx=".5" transform="rotate(90 14 5)"/><rect width="1" height="1" x="8" y="3" rx=".5" transform="rotate(90 8 3)"/><rect width="1" height="1" x="14" y="3" rx=".5" transform="rotate(90 14 3)"/><rect width="1" height="1" x="8" y="9" rx=".5" transform="rotate(90 8 9)"/><rect width="1" height="1" x="14" y="9" rx=".5" transform="rotate(90 14 9)"/><rect width="1" height="1" x="8" y="11" rx=".5" transform="rotate(90 8 11)"/><rect width="1" height="1" x="14" y="11" rx=".5" transform="rotate(90 14 11)"/><rect width="1" height="1" x="6" y="7" rx=".5" transform="rotate(90 6 7)"/><rect width="1" height="1" x="6" y="13" rx=".5" transform="rotate(90 6 13)"/><rect width="1" height="1" x="4" y="7" rx=".5" transform="rotate(90 4 7)"/><rect width="1" height="1" x="4" y="13" rx=".5" transform="rotate(90 4 13)"/><rect width="1" height="1" x="10" y="1" rx=".5" transform="rotate(90 10 1)"/><rect width="1" height="1" x="12" y="1" rx=".5" transform="rotate(90 12 1)"/><rect width="1" height="1" x="8" y="1" rx=".5" transform="rotate(90 8 1)"/><rect width="1" height="1" x="14" y="1" rx=".5" transform="rotate(90 14 1)"/><rect width="1" height="1" x="6" y="1" rx=".5" transform="rotate(90 6 1)"/><rect width="1" height="1" x="4" y="1" rx=".5" transform="rotate(90 4 1)"/></g>`),
		g.Group(children),
	)
}