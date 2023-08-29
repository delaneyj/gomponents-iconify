package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosAlbumsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M464 144v288H48V144h416m16-16H32v320h448V128z" fill="currentColor"/><path d="M72 96h368v16H72z" fill="currentColor"/><path d="M104 64h304v16H104z" fill="currentColor"/>`),
		g.Group(children),
	)
}