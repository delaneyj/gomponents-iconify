package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.627 3.002H11V2h1.627c.362 0 .71.144.967.402a1.392 1.392 0 0 1-.17 2.11l-.675.486h1.232V6H11v-.974l1.847-1.33a.385.385 0 0 0 .047-.583a.378.378 0 0 0-.267-.111ZM2.22 4.22a.75.75 0 0 1 1.06 0L6 6.94l2.72-2.72a.75.75 0 0 1 1.06 1.06L7.06 8l2.72 2.72a.75.75 0 1 1-1.06 1.06L6 9.06l-2.72 2.72a.75.75 0 0 1-1.06-1.06L4.94 8L2.22 5.28a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}