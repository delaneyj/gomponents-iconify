package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M14 1.75H1V.25h13v1.5Z" clip-rule="evenodd"/><rect width="1" height="1" x="8" y="10" rx=".5" transform="rotate(-180 8 10)"/><rect width="1" height="1" x="2" y="10" rx=".5" transform="rotate(-180 2 10)"/><rect width="1" height="1" x="8" y="12" rx=".5" transform="rotate(-180 8 12)"/><rect width="1" height="1" x="2" y="12" rx=".5" transform="rotate(-180 2 12)"/><rect width="1" height="1" x="8" y="8" rx=".5" transform="rotate(-180 8 8)"/><rect width="1" height="1" x="8" y="14" rx=".5" transform="rotate(-180 8 14)"/><rect width="1" height="1" x="2" y="8" rx=".5" transform="rotate(-180 2 8)"/><rect width="1" height="1" x="2" y="14" rx=".5" transform="rotate(-180 2 14)"/><rect width="1" height="1" x="10" y="8" rx=".5" transform="rotate(-180 10 8)"/><rect width="1" height="1" x="10" y="14" rx=".5" transform="rotate(-180 10 14)"/><rect width="1" height="1" x="12" y="8" rx=".5" transform="rotate(-180 12 8)"/><rect width="1" height="1" x="12" y="14" rx=".5" transform="rotate(-180 12 14)"/><rect width="1" height="1" x="6" y="8" rx=".5" transform="rotate(-180 6 8)"/><rect width="1" height="1" x="6" y="14" rx=".5" transform="rotate(-180 6 14)"/><rect width="1" height="1" x="4" y="8" rx=".5" transform="rotate(-180 4 8)"/><rect width="1" height="1" x="4" y="14" rx=".5" transform="rotate(-180 4 14)"/><rect width="1" height="1" x="8" y="6" rx=".5" transform="rotate(-180 8 6)"/><rect width="1" height="1" x="2" y="6" rx=".5" transform="rotate(-180 2 6)"/><rect width="1" height="1" x="8" y="4" rx=".5" transform="rotate(-180 8 4)"/><rect width="1" height="1" x="2" y="4" rx=".5" transform="rotate(-180 2 4)"/><rect width="1" height="1" x="14" y="10" rx=".5" transform="rotate(-180 14 10)"/><rect width="1" height="1" x="14" y="12" rx=".5" transform="rotate(-180 14 12)"/><rect width="1" height="1" x="14" y="8" rx=".5" transform="rotate(-180 14 8)"/><rect width="1" height="1" x="14" y="14" rx=".5" transform="rotate(-180 14 14)"/><rect width="1" height="1" x="14" y="6" rx=".5" transform="rotate(-180 14 6)"/><rect width="1" height="1" x="14" y="4" rx=".5" transform="rotate(-180 14 4)"/></g>`),
		g.Group(children),
	)
}