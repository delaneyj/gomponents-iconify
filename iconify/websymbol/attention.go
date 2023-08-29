package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attention(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1113 920q0 39-27.5 59.5T1018 1000H95q-40 0-67.5-21T0 920q0-30 18-61L479 55q33-55 78-55t76 55l462 805q18 32 18 60zM626 409V263H487v146q0 14 2 26.5t5.5 28T500 490l26 162h59l27-162q2-10 6-26t6-28.5t2-26.5zm0 387q0-29-20.5-49T556 727q-28 0-48.5 20T487 796t20.5 49.5T556 866q29 0 49.5-20.5T626 796z"/>`),
		g.Group(children),
	)
}