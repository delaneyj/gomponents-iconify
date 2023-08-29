package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlighterCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198.71 57.29A100 100 0 1 0 57.29 198.71A100 100 0 1 0 198.71 57.29ZM92 212.7V152a4 4 0 0 1 4-4h64a4 4 0 0 1 4 4v60.7a92.42 92.42 0 0 1-72 0Zm56-72.7h-40V98.47l40-20Zm45.05 53.05A92 92 0 0 1 172 208.83V152a12 12 0 0 0-12-12h-4V72a4 4 0 0 0-5.79-3.58l-48 24A4 4 0 0 0 100 96v44h-4a12 12 0 0 0-12 12v56.83a92 92 0 1 1 109.05-15.78Z"/>`),
		g.Group(children),
	)
}