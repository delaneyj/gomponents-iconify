package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageClass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4H2v4h20ZM6 7H4V5h2Zm14.6 7.8A2.088 2.088 0 0 0 19 14h-7.002a.998.998 0 0 0-.998.998v6.004a.998.998 0 0 0 .998.998H19a1.816 1.816 0 0 0 1.6-.8L23 18ZM11.998 12H19a4.075 4.075 0 0 1 3 1.39V10H2v4h7.184a2.993 2.993 0 0 1 2.814-2ZM4 13v-2h2v2Zm5 3H2v4h7Zm-5 3v-2h2v2Z"/>`),
		g.Group(children),
	)
}