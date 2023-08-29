package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22a6 6 0 0 1-5.14-2.92l-1.71 1a8 8 0 0 0 13.7 0l-1.71-1A6 6 0 0 1 16 22zM30 4h-4V0h-2v4h-4v2h4v4h2V6h4V4zm-18.5 7a2.5 2.5 0 1 0 2.5 2.5a2.48 2.48 0 0 0-2.5-2.5zm9 0a2.5 2.5 0 1 0 2.5 2.5a2.48 2.48 0 0 0-2.5-2.5z"/><path fill="currentColor" d="M27.61 13a12.33 12.33 0 0 1 .39 3A12 12 0 1 1 16 4V2a14 14 0 1 0 13.67 11Z"/>`),
		g.Group(children),
	)
}