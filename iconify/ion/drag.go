package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M0 144h512v32H0z" fill="currentColor"/><path d="M0 240h512v32H0z" fill="currentColor"/><path d="M0 336h512v32H0z" fill="currentColor"/>`),
		g.Group(children),
	)
}