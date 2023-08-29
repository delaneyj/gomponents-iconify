package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.83 9.17a1 1 0 1 0-1.42 1.42A2 2 0 0 1 18 12a2 2 0 0 1-.71 1.53a1 1 0 0 0-.13 1.41a1 1 0 0 0 1.41.12A4 4 0 0 0 20 12a4.06 4.06 0 0 0-1.17-2.83Zm-4.4-5.07a1 1 0 0 0-1 .12L8.65 8H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.65l4.73 3.78A1 1 0 0 0 14 20a.91.91 0 0 0 .43-.1A1 1 0 0 0 15 19V5a1 1 0 0 0-.57-.9ZM13 16.92l-3.38-2.7A1 1 0 0 0 9 14H6v-4h3a1 1 0 0 0 .62-.22L13 7.08Z"/>`),
		g.Group(children),
	)
}