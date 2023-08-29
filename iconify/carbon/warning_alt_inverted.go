package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningAltInverted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 17a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 16 17zM15 6h2v9h-2z"/><path fill="currentColor" d="M29.855 2.481a1.001 1.001 0 0 1 .032.98l-13 25a1 1 0 0 1-1.774 0l-13-25A1 1 0 0 1 3 2h26a1 1 0 0 1 .855.481ZM4.649 4.003L15.998 25.83h.004l11.35-21.826L27.348 4H4.651Z"/>`),
		g.Group(children),
	)
}