package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#C33" d="m67.96 38.21l28 48.5a3.749 3.749 0 0 1-3.3 5.4h-56c-2.03.03-3.7-1.6-3.73-3.63c-.01-.62.14-1.22.43-1.77l28-48.5a3.642 3.642 0 0 1 6.6 0z"/><path fill="#F44336" d="m64.6 43.01l25.5 44.2a3.378 3.378 0 0 1-3 4.9h-51a3.348 3.348 0 0 1-3.37-3.35c0-.54.13-1.07.37-1.55l25.5-44.2a3.272 3.272 0 0 1 4.33-1.67c.74.32 1.34.92 1.67 1.67z"/><path fill="#FF8A80" d="M60.72 47.85c-1.02 1.14-18.79 30.98-18.79 30.98s-1.52 2.29.25 3.68c1.78 1.4 3.43 0 4.19-1.14c.76-1.14 16.89-27.93 17.52-29.2a4.45 4.45 0 0 0-.63-4.32c-.89-.89-1.9-.89-2.54 0z"/>`),
		g.Group(children),
	)
}