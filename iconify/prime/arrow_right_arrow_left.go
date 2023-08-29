package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.75 16c0 .41-.34.75-.75.75H6.81l1.22 1.22c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22s-.38-.07-.53-.22l-2.5-2.5a.776.776 0 0 1-.16-.24a.707.707 0 0 1 0-.57c.04-.09.09-.17.16-.24l2.5-2.5c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-1.22 1.22H19c.41 0 .75.34.75.75ZM5 8.75h12.19l-1.22 1.22c-.29.29-.29.77 0 1.06c.15.15.34.22.53.22s.38-.07.53-.22l2.5-2.5c.07-.07.12-.15.16-.24c.08-.18.08-.39 0-.57a.776.776 0 0 0-.16-.24l-2.5-2.5c-.29-.29-.77-.29-1.06 0s-.29.77 0 1.06l1.22 1.22H5c-.41 0-.75.34-.75.75s.34.75.75.75Z"/>`),
		g.Group(children),
	)
}