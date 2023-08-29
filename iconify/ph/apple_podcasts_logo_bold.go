package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplePodcastsLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M161.36 130a40 40 0 1 0-66.72 0a28.35 28.35 0 0 0-4.52 4.51a27.86 27.86 0 0 0-5.42 23.67l12.75 56A27.87 27.87 0 0 0 124.77 236h6.46a27.87 27.87 0 0 0 27.32-21.79l12.75-56a27.86 27.86 0 0 0-5.42-23.67a28.35 28.35 0 0 0-4.52-4.54ZM128 92a16 16 0 1 1-16 16a16 16 0 0 1 16-16Zm19.89 60.88l-12.74 56a4 4 0 0 1-3.92 3.12h-6.46a4 4 0 0 1-3.92-3.12l-12.74-56a3.92 3.92 0 0 1 .77-3.37A4 4 0 0 1 112 148h32a4 4 0 0 1 3.15 1.51a3.92 3.92 0 0 1 .74 3.37ZM236 128a107.25 107.25 0 0 1-31.77 76.5a12 12 0 0 1-16.94-17a84 84 0 1 0-118.58 0a12 12 0 1 1-16.94 17A108 108 0 1 1 236 128Z"/>`),
		g.Group(children),
	)
}