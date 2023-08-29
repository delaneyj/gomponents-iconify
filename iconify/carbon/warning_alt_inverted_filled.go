package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningAltInvertedFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M16 20a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 20Zm-1.125-5h2.25V6h-2.25Z"/><path fill="currentColor" d="M27.35 4H4.65l-.001.003L15.998 25.83h.004l11.35-21.826ZM14.874 6h2.25v9h-2.25ZM16 20a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 20Z"/><path fill="currentColor" d="M29.855 2.481a1.001 1.001 0 0 1 .032.98l-13 25a1 1 0 0 1-1.774 0l-13-25A1 1 0 0 1 3 2h26a1 1 0 0 1 .855.481ZM4.649 4.003L15.998 25.83h.004l11.35-21.826L27.348 4H4.651Z"/>`),
		g.Group(children),
	)
}