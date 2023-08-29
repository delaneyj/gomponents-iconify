package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v6H8V8.5H4V11H2V5Zm2 1.5h4V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v1.5Zm8-3h6a3 3 0 0 1 3 3V9h-2V6.5a1 1 0 0 0-1-1h-6v-2Zm6 8V13h4v2h-1.062a7.974 7.974 0 0 1-2.19 4.563A5.984 5.984 0 0 0 21 20h1v2h-1a7.963 7.963 0 0 1-4-1.07A7.963 7.963 0 0 1 13 22h-1v-2h1c.796 0 1.556-.155 2.251-.437a8.013 8.013 0 0 1-1.48-2.134l-.43-.903l1.807-.858l.429.903A6.02 6.02 0 0 0 17 18.472A5.99 5.99 0 0 0 18.917 15H12v-2h4v-1.5h2ZM6 13v6a1 1 0 0 0 1 1h2.5v2H7a3 3 0 0 1-3-3v-6h2Z"/>`),
		g.Group(children),
	)
}