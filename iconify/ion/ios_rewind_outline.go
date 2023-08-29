package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosRewindOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M464 155v201.9L280.5 256 464 155m-224 1v200.4L64 256l176-100.2m16-27.8L32 256l224 128V260.8L480 384V128L256 251.2V128z" fill="currentColor"/>`),
		g.Group(children),
	)
}