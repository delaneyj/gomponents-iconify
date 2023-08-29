package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraseThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M30.415 16.513l-7.927-7.927a2.001 2.001 0 0 0-2.83 0L5.622 22.624a2.002 2.002 0 0 0 0 2.83L10.166 30h9.591l10.658-10.659a2.001 2.001 0 0 0 0-2.828zM18.929 28h-7.934l-3.96-3.962l6.312-6.312l7.928 7.928zm3.76-3.76l-7.928-7.928L21.074 10l7.927 7.927z" fill="currentColor"/><path d="M11 12H8V2h3a3.003 3.003 0 0 1 3 3v4a3.003 3.003 0 0 1-3 3zm-1-2h1a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-1z" fill="currentColor"/><path d="M4 2H0v2h4v2H1v2h3v2H0v2h4a2.003 2.003 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}