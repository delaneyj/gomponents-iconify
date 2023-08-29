package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="21" cy="26" r="2" fill="currentColor"/><circle cx="21" cy="6" r="2" fill="currentColor"/><circle cx="4" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M28 12a3.996 3.996 0 0 0-3.858 3h-4.284a3.966 3.966 0 0 0-5.491-2.643l-3.177-3.97A3.963 3.963 0 0 0 12 6a4 4 0 1 0-4 4a3.96 3.96 0 0 0 1.634-.357l3.176 3.97a3.924 3.924 0 0 0 0 4.774l-3.176 3.97A3.96 3.96 0 0 0 8 22a4 4 0 1 0 4 4a3.962 3.962 0 0 0-.81-2.387l3.176-3.97A3.966 3.966 0 0 0 19.858 17h4.284A3.993 3.993 0 1 0 28 12ZM6 6a2 2 0 1 1 2 2a2.002 2.002 0 0 1-2-2Zm2 22a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm8-10a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm12 0a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}