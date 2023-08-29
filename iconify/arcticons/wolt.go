package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.983 18.259l-4.896 11.482l-.846-11.482l-4.895 11.482l-.846-11.482m13.798 11.482h0c-1.585 0-2.644-1.285-2.364-2.87l.329-1.866c.28-1.586 1.791-2.871 3.377-2.871h0c1.585 0 2.644 1.285 2.364 2.87l-.329 1.867c-.28 1.585-1.791 2.87-3.377 2.87Zm7.747-11.482l-1.772 10.047c-.14.793.39 1.436 1.182 1.436h.431m4.524-9.976l-1.505 8.54c-.14.793.39 1.435 1.182 1.435h.43m-2.031-7.607H37.5"/>`),
		g.Group(children),
	)
}