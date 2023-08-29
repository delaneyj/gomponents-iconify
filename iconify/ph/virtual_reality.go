package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualReality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m123.41 99l-26 64a8 8 0 0 1-14.82 0l-26-64a8 8 0 1 1 14.82-6L90 138.74L108.59 93a8 8 0 1 1 14.82 6ZM256 128a80.09 80.09 0 0 1-80 80H80a80 80 0 0 1 0-160h96a80.09 80.09 0 0 1 80 80Zm-16 0a64.07 64.07 0 0 0-64-64H80a64 64 0 0 0 0 128h96a64.07 64.07 0 0 0 64-64Zm-59.16 10.35L191 156a8 8 0 0 1-13.9 7.94l-11.44-20c-.53 0-1.07.05-1.61.05H152v16a8 8 0 0 1-16 0V96a8 8 0 0 1 8-8h20a28 28 0 0 1 16.84 50.35ZM152 128h12a12 12 0 0 0 0-24h-12Z"/>`),
		g.Group(children),
	)
}