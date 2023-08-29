package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTransferUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 16h480v32H16zm100.687 196.687l22.626 22.626L232 142.627V496h32V142.627l92.687 92.686l22.626-22.626L248 81.373L116.687 212.687z"/>`),
		g.Group(children),
	)
}