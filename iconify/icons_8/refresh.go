package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4c-5.11 0-9.383 3.16-11.125 7.625l1.844.75C8.176 8.64 11.71 6 16 6c3.24 0 6.134 1.59 7.938 4H20v2h7V5h-2v3.094A11.928 11.928 0 0 0 16 4zm9.28 15.625C23.824 23.36 20.29 26 16 26c-3.276 0-6.157-1.612-7.97-4H12v-2H5v7h2v-3.094C9.19 26.386 12.395 28 16 28c5.11 0 9.383-3.16 11.125-7.625l-1.844-.75z"/>`),
		g.Group(children),
	)
}