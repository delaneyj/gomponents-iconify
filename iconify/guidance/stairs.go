package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 24v-.149a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12V6.5h7v-.648A7.5 7.5 0 0 0 16.606.87L16.5.749V.5H24"/>`),
		g.Group(children),
	)
}