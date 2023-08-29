package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.998 12.127v7.896h4.495l6.73 5.526l.003-18.95l-6.73 5.527H4.998zm13.808-.908a1.002 1.002 0 0 0-1.415 0c-.39.392-.39 1.025.003 1.417v-.002A4.741 4.741 0 0 1 18.789 16c0 1.317-.53 2.498-1.393 3.362a.996.996 0 0 0-.002 1.415a.998.998 0 0 0 1.415 0a6.74 6.74 0 0 0 1.98-4.776c0-1.864-.76-3.56-1.982-4.78z"/>`),
		g.Group(children),
	)
}