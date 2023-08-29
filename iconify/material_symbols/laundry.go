package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laundry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 13.975v-3L4.125 12l-3-5.2l6.6-3.8H9q.4 1.2.95 2.1T12 6q1.5 0 2.05-.9T15 3h1.275l6.575 3.825L19.875 12L18 10.975v4.8l-1.575 1.375q-.175.15-.4.238t-.45.087q-.15 0-.35-.1t-.4-.25l-2.65-2.275q-.8-.7-1.8-1.025T8.35 13.5q-.6 0-1.188.113T6 13.975Zm-1.35 5.4l-1.3-1.525L5.525 16q.575-.5 1.313-.763t1.537-.262q.8 0 1.525.263t1.3.762l2.9 2.475q.3.25.713.388t.837.137q.45 0 .838-.125t.687-.4L19.35 16.6l1.3 1.55L18.475 20q-.575.5-1.3.75T15.65 21q-.8 0-1.538-.25T12.8 20l-2.9-2.475q-.3-.25-.688-.388T8.376 17q-.425 0-.838.138t-.712.387l-2.175 1.85Z"/>`),
		g.Group(children),
	)
}