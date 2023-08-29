package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.44 9.33h-1.1c0-.97.01-1.95 0-2.92a4.343 4.343 0 0 0-7.98-2.37c-.36.53.51 1.03.87.5a3.365 3.365 0 0 1 5.23-.39c1.04 1.11.88 2.57.88 3.96v1.22H6.56a2.5 2.5 0 0 0-2.5 2.5v7.61a2.507 2.507 0 0 0 2.5 2.5h10.88a2.507 2.507 0 0 0 2.5-2.5v-7.61a2.5 2.5 0 0 0-2.5-2.5Zm1.5 10.11a1.511 1.511 0 0 1-1.5 1.5H6.56a1.511 1.511 0 0 1-1.5-1.5v-7.61a1.5 1.5 0 0 1 1.5-1.5h10.88a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M13 14.95a.984.984 0 0 1-.5.86v1.5a.5.5 0 0 1-1 0v-1.5a.984.984 0 0 1-.5-.86a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}