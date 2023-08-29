package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemCoordinate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.272 1v2.65l6.613 3.89v7.745l2.465 1.295l-.93 1.77l-2.507-1.317l-5.641 3.317V23h-2v-2.65l-5.69-3.346l-3.12 1.304l-.77-1.845l2.967-1.24V7.539l6.613-3.889V1h2Zm-1 4.382l-4.64 2.73l4.64 2.728l4.64-2.729l-4.64-2.729Zm5.613 4.477l-4.613 2.713v5.458l4.613-2.713V9.859Zm-6.613 8.17v-5.457L6.66 9.859v5.458l4.613 2.713Z"/>`),
		g.Group(children),
	)
}