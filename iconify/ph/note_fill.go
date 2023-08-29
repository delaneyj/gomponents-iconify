package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h108.69a15.92 15.92 0 0 0 11.31-4.69L219.31 168a15.92 15.92 0 0 0 4.69-11.31V48a16 16 0 0 0-16-16ZM96 88h64a8 8 0 0 1 0 16H96a8 8 0 0 1 0-16Zm32 80H96a8 8 0 0 1 0-16h32a8 8 0 0 1 0 16Zm-32-32a8 8 0 0 1 0-16h64a8 8 0 0 1 0 16Zm64 68.69V160h44.7Z"/>`),
		g.Group(children),
	)
}