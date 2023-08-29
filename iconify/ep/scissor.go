package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m512.064 578.368l-106.88 152.768a160 160 0 1 1-23.36-78.208L472.96 522.56L196.864 128.256a32 32 0 1 1 52.48-36.736l393.024 561.344a160 160 0 1 1-23.36 78.208l-106.88-152.704zm54.4-189.248l208.384-297.6a32 32 0 0 1 52.48 36.736l-221.76 316.672l-39.04-55.808zm-376.32 425.856a96 96 0 1 0 110.144-157.248a96 96 0 0 0-110.08 157.248zm643.84 0a96 96 0 1 0-110.08-157.248a96 96 0 0 0 110.08 157.248z"/>`),
		g.Group(children),
	)
}