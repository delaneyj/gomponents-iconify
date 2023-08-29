package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSoundOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.4 16L14 14.6l2.6-2.6L14 9.4L15.4 8l2.6 2.6L20.6 8L22 9.4L19.4 12l2.6 2.6l-1.4 1.4l-2.6-2.6l-2.6 2.6ZM3 15V9h4l5-5v16l-5-5H3Zm7-6.15L7.85 11H5v2h2.85L10 15.15v-6.3ZM7.5 12Z"/>`),
		g.Group(children),
	)
}