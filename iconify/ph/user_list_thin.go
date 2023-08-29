package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserListThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 80a4 4 0 0 1 4-4h96a4 4 0 0 1 0 8h-96a4 4 0 0 1-4-4Zm100 44h-96a4 4 0 0 0 0 8h96a4 4 0 0 0 0-8Zm0 48h-72a4 4 0 0 0 0 8h72a4 4 0 0 0 0-8Zm-100.13 19a4 4 0 0 1-2.87 4.87a3.87 3.87 0 0 1-1 .13a4 4 0 0 1-3.87-3c-6.71-26.08-32-45-60.13-45s-53.41 18.92-60.13 45a4 4 0 1 1-7.74-2c5.92-23 24.57-41.14 47.52-48a44 44 0 1 1 40.7 0c22.95 6.86 41.65 25 47.52 48ZM80 140a36 36 0 1 0-36-36a36 36 0 0 0 36 36Z"/>`),
		g.Group(children),
	)
}