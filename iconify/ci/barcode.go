package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h3v16H2V4Zm4 0h1v16H6V4Zm5 0H9v16h2V4Zm1 0h2v16h-2V4Zm3 0h1v16h-1V4Zm5 0h-3v16h3V4Zm1 0h1v16h-1V4Z"/>`),
		g.Group(children),
	)
}