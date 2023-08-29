package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualRealityLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m121.56 98.26l-26 64a6 6 0 0 1-11.12 0l-26-64a6 6 0 0 1 11.12-4.52L90 144.06l20.44-50.32a6 6 0 0 1 11.12 4.52ZM254 128a78.09 78.09 0 0 1-78 78H80a78 78 0 0 1 0-156h96a78.09 78.09 0 0 1 78 78Zm-12 0a66.08 66.08 0 0 0-66-66H80a66 66 0 0 0 0 132h96a66.08 66.08 0 0 0 66-66Zm-63.8 9.76l11 19.26a6 6 0 0 1-10.42 6l-12.07-21.12a27.06 27.06 0 0 1-2.71.1h-14v18a6 6 0 0 1-12 0V96a6 6 0 0 1 6-6h20a26 26 0 0 1 14.2 47.76ZM150 130h14a14 14 0 0 0 0-28h-14Z"/>`),
		g.Group(children),
	)
}