package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 5C5.935 5 1 9.935 1 16s4.935 11 11 11h19v-2H18.305C21.139 23.008 23 19.72 23 16c0-6.065-4.935-11-11-11zm0 2c4.962 0 9 4.037 9 9s-4.038 9-9 9s-9-4.037-9-9s4.038-9 9-9zm0 5c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4s-1.794-4-4-4zm0 2c1.103 0 2 .897 2 2s-.897 2-2 2s-2-.897-2-2s.897-2 2-2z"/>`),
		g.Group(children),
	)
}