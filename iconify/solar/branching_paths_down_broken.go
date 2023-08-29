package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchingPathsDownBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12.75 6.5a.75.75 0 0 0-1.5 0h1.5Zm5.25 10l.53.53a.75.75 0 0 0 0-1.06l-.53.53Zm-2.03.97a.75.75 0 1 0 1.06 1.06l-1.06-1.06Zm1.06-3a.75.75 0 1 0-1.06 1.06l1.06-1.06ZM11.25 6.5v6h1.5v-6h-1.5ZM16 17.25h2v-1.5h-2v1.5Zm1.47-1.28l-1.5 1.5l1.06 1.06l1.5-1.5l-1.06-1.06Zm1.06 0l-1.5-1.5l-1.06 1.06l1.5 1.5l1.06-1.06Zm-7.28-3.47A4.75 4.75 0 0 0 16 17.25v-1.5a3.25 3.25 0 0 1-3.25-3.25h-1.5Z"/><path fill="currentColor" d="M12.75 6.5a.75.75 0 0 0-1.5 0h1.5ZM6 16.5l-.53-.53a.75.75 0 0 0 0 1.06L6 16.5Zm.97 2.03a.75.75 0 0 0 1.06-1.06l-1.06 1.06Zm1.06-3a.75.75 0 1 0-1.06-1.06l1.06 1.06Zm3.22-9.03v6h1.5v-6h-1.5ZM8 15.75H6v1.5h2v-1.5Zm-2.53 1.28l1.5 1.5l1.06-1.06l-1.5-1.5l-1.06 1.06Zm1.06 0l1.5-1.5l-1.06-1.06l-1.5 1.5l1.06 1.06Zm4.72-4.53A3.25 3.25 0 0 1 8 15.75v1.5a4.75 4.75 0 0 0 4.75-4.75h-1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464c.974.974 1.3 2.343 1.41 4.536"/></g>`),
		g.Group(children),
	)
}