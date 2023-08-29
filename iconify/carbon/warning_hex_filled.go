package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningHexFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M14.875 8h2.25v11h-2.25ZM16 25a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 25Z"/><path fill="currentColor" d="M30.85 15.449L23.888 3.532A1.069 1.069 0 0 0 22.964 3H9.036a1.069 1.069 0 0 0-.923.532L1.15 15.45a1.093 1.093 0 0 0 0 1.102l6.964 11.917a1.069 1.069 0 0 0 .923.532h13.928a1.069 1.069 0 0 0 .923-.532L30.85 16.55a1.093 1.093 0 0 0 0-1.102ZM14.876 8h2.25v11h-2.25ZM16 25a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 25Z"/>`),
		g.Group(children),
	)
}