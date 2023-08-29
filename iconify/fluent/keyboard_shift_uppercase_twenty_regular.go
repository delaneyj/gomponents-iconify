package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardShiftUppercaseTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.143 2.53a1.5 1.5 0 0 0-2.288 0L2.24 10.333A1 1 0 0 0 3 11.98h3v3.017a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V11.98h2.996a1 1 0 0 0 .763-1.647L11.143 2.53Zm-1.525.647a.5.5 0 0 1 .763 0l6.616 7.803H13.5a.5.5 0 0 0-.5.5v3.517H7V11.48a.5.5 0 0 0-.5-.5H3.001l6.617-7.803ZM6.5 17a.5.5 0 1 0 0 1h7a.5.5 0 0 0 0-1h-7Z"/>`),
		g.Group(children),
	)
}