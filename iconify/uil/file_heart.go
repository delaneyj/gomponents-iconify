package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.94Zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-6-7.66A2.92 2.92 0 0 0 8.57 16l2.72 2.72a1 1 0 0 0 1.42 0L15.43 16A2.92 2.92 0 0 0 12 11.34Zm2 1.93a.92.92 0 0 1 0 1.3l-2 2l-2-2a.92.92 0 0 1 0-1.3a.92.92 0 0 1 1.3 0a1 1 0 0 0 1.42 0a.92.92 0 0 1 1.28 0Z"/>`),
		g.Group(children),
	)
}