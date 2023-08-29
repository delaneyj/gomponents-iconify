package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 5.835A10.485 10.485 0 0 0 12 1.5c-5.427 0-9.89 4.115-10.443 9.396l-.104.994l1.99.209l.103-.995A8.501 8.501 0 0 1 19.212 7.5H15.5v2h7v-7h-2v3.335Zm.057 6.066l-.104.995A8.501 8.501 0 0 1 4.787 16.5H8.5v-2h-7v7h2v-3.335A10.485 10.485 0 0 0 12 22.5c5.426 0 9.89-4.115 10.442-9.396l.104-.994l-1.989-.209Z"/>`),
		g.Group(children),
	)
}