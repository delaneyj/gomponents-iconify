package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachmentAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M203 112h32v245q0 49-34.5 83.5t-83 34.5t-83-34.5T0 357V91q0-36 25-61T85.5 5T146 30t25 61v224q0 22-16 37.5T117 368t-37.5-15.5T64 315V112h32v203q0 8 6.5 14.5t15 6.5t15-6.5T139 315V91q0-22-16-38T85 37T47.5 53T32 91v266q0 36 25 61t60.5 25t60.5-25t25-61V112z"/>`),
		g.Group(children),
	)
}