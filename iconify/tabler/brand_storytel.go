package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandStorytel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.103 22c2.292-2.933 16.825-2.43 16.825-11.538C20.928 4.164 15.954 2 12.477 2C9 2 3 5.036 3 13.241C3 19.615 4.103 22 4.103 22z"/>`),
		g.Group(children),
	)
}