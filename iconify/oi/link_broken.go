package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2 0v1H1v1h2V0H2zm3.88.03a1.9 1.9 0 0 0-.53.09c-.27.1-.53.25-.75.47l-.44.44a.5.5 0 1 0 .69.69l.44-.44c.11-.11.24-.17.38-.22c.35-.12.78-.07 1.06.22c.39.39.39 1.04 0 1.44l-1.5 1.5a.5.5 0 1 0 .69.69l1.5-1.5A1.98 1.98 0 0 0 6.45.07C6.27.03 6.07.03 5.89.04zM2.29 2.94a.5.5 0 0 0-.19.16L.6 4.6a1.98 1.98 0 0 0 0 2.81c.56.56 1.36.72 2.06.47c.27-.1.53-.25.75-.47l.44-.44a.5.5 0 1 0-.69-.69l-.44.44c-.11.11-.24.17-.38.22c-.35.12-.78.07-1.06-.22c-.39-.39-.39-1.04 0-1.44l1.5-1.5a.5.5 0 0 0-.44-.84a.5.5 0 0 0-.06 0zM5.01 6v2h1V7h1V6h-2z"/>`),
		g.Group(children),
	)
}