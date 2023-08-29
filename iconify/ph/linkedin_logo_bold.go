package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 20H40a20 20 0 0 0-20 20v176a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V40a20 20 0 0 0-20-20Zm-4 192H44V44h168Zm-100-36v-52a12 12 0 0 1 21.43-7.41A40 40 0 0 1 192 152v24a12 12 0 0 1-24 0v-24a16 16 0 0 0-32 0v24a12 12 0 0 1-24 0Zm-16-52v52a12 12 0 0 1-24 0v-52a12 12 0 0 1 24 0ZM68 80a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}