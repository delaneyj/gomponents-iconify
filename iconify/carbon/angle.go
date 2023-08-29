package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M9 24a3.51 3.51 0 0 0-.88-1.86l9.53-16.58l-1.73-1l-9.57 16.56A3.06 3.06 0 0 0 5.5 21a3.5 3.5 0 1 0 3.15 5H28v-2zm-3.5 2A1.5 1.5 0 1 1 7 24.5A1.5 1.5 0 0 1 5.5 26z" fill="currentColor"/><path d="M22 21h2a13 13 0 0 0-5.42-10.56l-1.16 1.62A11 11 0 0 1 22 21z" fill="currentColor"/>`),
		g.Group(children),
	)
}