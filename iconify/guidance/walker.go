package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10.004 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-3.5-.75-5v-.5l-5-2.5v-.25l1.803-5.228A3 3 0 0 1 7.143 6.5h1.36l.519 4.588a3 3 0 0 0 2.941 2.412H12m-6.496 4c-.21 0-.466.133-.737.344C2.962 19.24 2.503 21.718 2.503 24m8.996-.2a26 26 0 0 0 2-10v-.3h6v.3a26 26 0 0 0 2 10m-8.878-3.3h7.756M8.354 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}