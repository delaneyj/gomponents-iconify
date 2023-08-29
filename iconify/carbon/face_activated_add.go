package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceActivatedAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 4h-4V0h-2v4h-4v2h4v4h2V6h4V4zm-18.5 7a2.5 2.5 0 1 0 2.5 2.5a2.48 2.48 0 0 0-2.5-2.5zm9 0a2.5 2.5 0 1 0 2.5 2.5a2.48 2.48 0 0 0-2.5-2.5zM9 20a8.13 8.13 0 0 0 14 0z"/><path fill="currentColor" d="M27.82 14A12 12 0 1 1 16 4V2a14 14 0 1 0 14 14a14.71 14.71 0 0 0-.16-2Z"/>`),
		g.Group(children),
	)
}