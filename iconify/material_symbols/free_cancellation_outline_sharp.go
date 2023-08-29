package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeCancellationOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.55 22.5L13 18.95l1.4-1.4l2.125 2.125l4.25-4.25l1.4 1.425l-5.625 5.65ZM7.4 17L6 15.6L7.6 14L6 12.4L7.4 11L9 12.6l1.6-1.6l1.4 1.4l-1.6 1.6l1.6 1.6l-1.4 1.4L9 15.4L7.4 17ZM3 22V4h3V2h2v2h8V2h2v2h3v8.35l-2 2.025V10H5v10h6.25l1.975 2H3ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}