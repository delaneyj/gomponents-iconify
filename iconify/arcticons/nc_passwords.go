package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NcPasswords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.5 21.523V15a10.5 10.5 0 0 0-21 0v6.523a13.5 13.5 0 1 0 21 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 16.5a13.422 13.422 0 0 1 6 1.42V15a6 6 0 0 0-12 0v2.92a13.421 13.421 0 0 1 6-1.42ZM27.5 28a3.5 3.5 0 1 0-5 3.15V36h3v-4.85a3.491 3.491 0 0 0 2-3.15Z"/>`),
		g.Group(children),
	)
}