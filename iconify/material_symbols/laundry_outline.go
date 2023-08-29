package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaundryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.125 12l-3-5.2l6.6-3.8H9q.4 1.2.95 2.1T12 6q1.5 0 2.05-.9T15 3h1.275l6.575 3.825L19.875 12L18 10.975v4.8l-1.575 1.375q-.075.05-.2.125T16 17.4V7.575L19.125 9.3l1-1.75L16.3 5.325q-.6 1.225-1.763 1.95T12 8q-1.375 0-2.538-.725T7.7 5.325L3.85 7.55L4.875 9.3L8 7.575V13.5q-.525.05-1.025.163T6 13.975v-3L4.125 12Zm.525 7.375l-1.3-1.525L5.525 16q.575-.5 1.313-.763t1.537-.262q.8 0 1.525.263t1.3.762l2.9 2.475q.3.25.713.388t.837.137q.45 0 .838-.125t.687-.4L19.35 16.6l1.3 1.55L18.475 20q-.575.5-1.3.75T15.65 21q-.8 0-1.538-.25T12.8 20l-2.9-2.475q-.3-.25-.688-.388T8.376 17q-.425 0-.838.138t-.712.387l-2.175 1.85ZM12 8Z"/>`),
		g.Group(children),
	)
}