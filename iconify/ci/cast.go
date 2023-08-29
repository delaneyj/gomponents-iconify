package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21h-7v-2h7V5H3v3H1V5a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Zm-9 0h-2a8.94 8.94 0 0 0-2.639-6.361A8.939 8.939 0 0 0 1 12v-2a10.929 10.929 0 0 1 7.776 3.22A10.927 10.927 0 0 1 12 21Zm-4 0H6a4.968 4.968 0 0 0-1.466-3.534A4.966 4.966 0 0 0 1 16v-2a6.955 6.955 0 0 1 4.951 2.049A6.956 6.956 0 0 1 8 21Zm-4 0H1v-3a3 3 0 0 1 3 3Z"/>`),
		g.Group(children),
	)
}