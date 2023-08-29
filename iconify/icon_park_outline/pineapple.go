package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pineapple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20.5 14.5C20 14 20 13.15 20 11c0-3.866 4-7 4-7s3 3.134 3 7c0 1.922 0 3-.5 3.5"/><path d="M17 17c-1.21-2.918-4.138-6.567-6-8c5.411-1.014 9.774 2.3 12 5m7 2c1.452-2.697 4.304-5.695 6-7c-5.643-1.058-9.871 2.261-12 5"/><path d="M16.725 17.374A38.81 38.81 0 0 0 15 19c-4.19 4.19-4.898 11.864-4.997 15.28c-.034 1.15-.026 2.322.409 3.387C10.879 38.81 11.875 40.406 14 42c4 3 16 3 20 0c2.116-1.587 3.112-3.175 3.581-4.318c.441-1.074.45-2.258.415-3.419c-.1-3.425-.812-11.08-4.996-15.263c-.6-.6-1.176-1.14-1.726-1.626c-4.258-3.767-10.291-3.767-14.55 0ZM32 18L11 29m5-11l21 11m-2 12L11 27m26 0L13 41"/></g>`),
		g.Group(children),
	)
}