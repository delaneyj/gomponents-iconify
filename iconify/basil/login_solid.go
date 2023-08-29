package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoginSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 19a.75.75 0 0 1 .75-.75h6a.25.25 0 0 0 .25-.25V6a.25.25 0 0 0-.25-.25h-6a.75.75 0 0 1 0-1.5h6c.966 0 1.75.784 1.75 1.75v12A1.75 1.75 0 0 1 18 19.75h-6a.75.75 0 0 1-.75-.75Z"/><path fill="currentColor" d="M3.5 13.115a1 1 0 0 0 1 1h4.856c.023.356.052.71.086 1.066l.03.305a.718.718 0 0 0 1.025.578a16.844 16.844 0 0 0 4.884-3.539l.03-.031a.721.721 0 0 0 0-.998l-.03-.031a16.842 16.842 0 0 0-4.884-3.539a.718.718 0 0 0-1.025.578l-.03.305c-.034.355-.063.71-.086 1.066H4.5a1 1 0 0 0-1 1v2.24Z"/>`),
		g.Group(children),
	)
}