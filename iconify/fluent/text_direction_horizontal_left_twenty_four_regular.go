package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionHorizontalLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m20.947 12.97l-3.753-9.496a.75.75 0 0 0-1.345-.104l-.05.105l-3.747 9.5a.75.75 0 0 0 1.352.643l.043-.092l.997-2.526h4.113l.995 2.52a.75.75 0 0 0 .876.454l.098-.031a.75.75 0 0 0 .452-.876l-.03-.098Zm-4.45-7.178L17.964 9.5h-2.928l1.461-3.708ZM12 7.75a.75.75 0 0 1-.75.75H5.56l.72.72a.75.75 0 1 1-1.06 1.06l-2-2a.75.75 0 0 1 0-1.06l2-2a.75.75 0 0 1 1.06 1.06L5.56 7h5.69a.75.75 0 0 1 .75.75Zm9 9.5a.75.75 0 0 1-.75.75H5.56l.72.72a.75.75 0 1 1-1.06 1.06l-2-2a.75.75 0 0 1 0-1.06l2-2a.75.75 0 0 1 1.06 1.06l-.72.72h14.69a.75.75 0 0 1 .75.75Z"/>`),
		g.Group(children),
	)
}