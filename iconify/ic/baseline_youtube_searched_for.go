package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineYoutubeSearchedFor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.01 14h-.8l-.27-.27a6.451 6.451 0 0 0 1.57-4.23c0-3.59-2.91-6.5-6.5-6.5s-6.5 3-6.5 6.5H2l3.84 4l4.16-4H6.51a4.5 4.5 0 0 1 9 0a4.507 4.507 0 0 1-6.32 4.12L7.71 15.1a6.474 6.474 0 0 0 7.52-.67l.27.27v.79l5.01 4.99L22 19l-4.99-5z"/>`),
		g.Group(children),
	)
}