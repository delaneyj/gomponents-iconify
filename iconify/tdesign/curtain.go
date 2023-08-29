package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Curtain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 7.465V20h2.914L6.093 9.999L6 10a3.982 3.982 0 0 1-2-.535Zm4.053-.032L8.921 20h6.158l.868-10.567A4.02 4.02 0 0 1 15 8.646A3.99 3.99 0 0 1 12 10a3.99 3.99 0 0 1-3-1.354a4.02 4.02 0 0 1-.947.787ZM8 6h2a2 2 0 1 0 4 0h2a2.001 2.001 0 0 0 2 2a2 2 0 0 0 2-2V4H4v2a2 2 0 1 0 4 0Zm12 3.465a3.981 3.981 0 0 1-2.093.534L17.085 20H20V9.465Z"/>`),
		g.Group(children),
	)
}