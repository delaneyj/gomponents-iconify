package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.43 4.1a1 1 0 0 0-1 .12L6.65 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.65l4.73 3.78A1 1 0 0 0 12 20a.91.91 0 0 0 .43-.1A1 1 0 0 0 13 19V5a1 1 0 0 0-.57-.9ZM11 16.92l-3.38-2.7A1 1 0 0 0 7 14H4v-4h3a1 1 0 0 0 .62-.22L11 7.08Zm4.14-12.83a1 1 0 1 0-.28 2a6 6 0 0 1 0 11.86a1 1 0 0 0 .14 2h.14a8 8 0 0 0 0-15.82Zm-.46 9.78a1 1 0 0 0 .32 2a1.13 1.13 0 0 0 .32-.05a4 4 0 0 0 0-7.54a1 1 0 0 0-.64 1.9a2 2 0 0 1 0 3.74Z"/>`),
		g.Group(children),
	)
}