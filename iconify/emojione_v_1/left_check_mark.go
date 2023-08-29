package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftCheckMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#37b34a" d="M56.734 5.081c-8.437 7.302-15.575 14.253-22.11 23.322c-2.882 4-6.087 8.708-8.182 13.153c-1.196 2.357-3.352 6.04-4.087 9.581c-4.02-3.74-8.338-7.985-12.756-11.31c-3.149-2.369-12.219 2.461-8.527 5.239c6.617 4.977 12.12 11.176 18.556 16.375c2.692 2.172 8.658-2.545 10.06-4.524c4.602-6.52 5.231-14.49 8.585-21.602c5.121-10.877 14.203-19.812 23.17-27.571c5.941-5.541-.195-6.563-4.7-2.663"/>`),
		g.Group(children),
	)
}