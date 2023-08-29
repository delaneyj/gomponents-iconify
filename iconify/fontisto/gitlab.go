package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.906 24L.403 14.723a1.073 1.073 0 0 1-.351-.497l-.002-.008a.926.926 0 0 1 .002-.609l-.002.007l1.463-4.437zM5.293.354l2.874 8.823H1.512L4.335.354a.517.517 0 0 1 .49-.353h.015h-.001L4.865 0c.212 0 .388.151.427.351v.003zm2.874 8.823h9.479L12.907 24zm17.595 4.436a.926.926 0 0 1-.002.609l.002-.007a1.074 1.074 0 0 1-.351.503l-.002.002L12.906 24L24.3 9.177zM21.477.354L24.3 9.177h-6.655L20.519.354a.436.436 0 0 1 .455-.353h-.001h.014c.227 0 .419.146.489.349l.001.004z"/>`),
		g.Group(children),
	)
}