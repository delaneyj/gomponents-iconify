package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parenright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 981h81c33-76 59-158 76-247s25-180 25-273c0-87-7-170-21-251S126 59 101 0H25c26 62 46 134 61 216c14 81 22 166 22 254c0 95-10 187-29 276S33 914 0 981z"/>`),
		g.Group(children),
	)
}