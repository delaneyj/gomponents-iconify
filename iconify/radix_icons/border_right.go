package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.25 1v13h1.5V1h-1.5Z" clip-rule="evenodd"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 5 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 5 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 3 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 3 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 5)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 5)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 3)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 3)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 9)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 9)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 11)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 11)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 9 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 9 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 11 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 11 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 5 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 3 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 9 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 11 1)"/></g>`),
		g.Group(children),
	)
}