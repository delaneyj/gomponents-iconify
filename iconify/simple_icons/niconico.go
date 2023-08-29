package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Niconico(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.479 7.534v12.128A2.021 2.021 0 0 0 2.5 21.683h2.389l1.323 2.095a.478.478 0 0 0 .404.22a.478.478 0 0 0 .441-.22l1.323-2.095h6.983l1.323 2.095a.478.478 0 0 0 .44.22c.185 0 .332-.073.405-.22l1.323-2.095H21.5a2.021 2.021 0 0 0 2.022-2.021V7.534A2.021 2.021 0 0 0 21.5 5.549h-7.68l4.446-4.447L17.164 0l-5.146 5.145L6.8 0L5.697 1.103l4.41 4.41h-7.57A2.021 2.021 0 0 0 .479 7.57z"/>`),
		g.Group(children),
	)
}