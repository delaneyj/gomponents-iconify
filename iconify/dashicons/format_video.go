package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 1h16c.55 0 1 .45 1 1v16l-18-.02V2c0-.55.45-1 1-1zm4 1L4 5h1l2-3H6zm4 0H9L7 5h1zm3 0h-1l-2 3h1zm3 0h-1l-2 3h1zm1 14V6H3v10h14zM8 7l6 4l-6 4V7z"/>`),
		g.Group(children),
	)
}