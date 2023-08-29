package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 28c-3.593 0-6.967-1.59-9.257-4.363l1.542-1.274A9.975 9.975 0 0 0 18 26c5.514 0 10-4.486 10-10S23.514 6 18 6a9.975 9.975 0 0 0-7.715 3.637L8.743 8.363A11.969 11.969 0 0 1 18 4c6.617 0 12 5.383 12 12s-5.383 12-12 12Z"/><path fill="currentColor" d="m23 16l-7-7l-1.414 1.414L19.172 15H2v2h17.172l-4.586 4.586L16 23l7-7z"/>`),
		g.Group(children),
	)
}