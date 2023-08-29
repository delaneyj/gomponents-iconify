package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chroma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<ellipse cx="170.667" cy="81.92" fill="#FFDE2D" rx="85.333" ry="81.92"/><ellipse cx="85.333" cy="81.92" fill="#327EFF" rx="85.333" ry="81.92"/><path fill="#FF6446" d="M170.667 81.92c0 45.243-38.206 81.92-85.334 81.92V81.92h85.334Zm-85.334 0C85.333 36.677 123.538 0 170.667 0v81.92H85.333Z"/>`),
		g.Group(children),
	)
}