package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m208.49 64.47l-144 144a12 12 0 1 1-17-17l144-144a12 12 0 0 1 17 17Zm-160.77 39.8A40 40 0 1 1 76 116a39.72 39.72 0 0 1-28.28-11.73ZM60 76a16 16 0 1 0 4.69-11.31A15.87 15.87 0 0 0 60 76Zm160 104a40 40 0 1 1-11.72-28.29A39.71 39.71 0 0 1 220 180Zm-24 0a15.87 15.87 0 0 0-4.69-11.32A16 16 0 1 0 196 180Z"/>`),
		g.Group(children),
	)
}