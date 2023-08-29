package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wurdian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.172 12.387V25.18c0 2.36-1.863 4.347-4.347 4.347c-2.36 0-4.347-1.863-4.347-4.347V12.387h-7.204V25.18c0 2.36-1.863 4.347-4.347 4.347s-4.347-1.863-4.347-4.347V12.387H4.5v13.165c0 5.59 4.844 10.061 10.682 10.061A11.12 11.12 0 0 0 24 31.266c1.987 2.608 5.093 4.347 8.819 4.347c5.961 0 10.681-4.471 10.681-10.06V12.387h-7.328Z"/>`),
		g.Group(children),
	)
}