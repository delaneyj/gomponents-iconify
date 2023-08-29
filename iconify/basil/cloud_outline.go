package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.932 6.208a3.91 3.91 0 0 0-3.524 2.214a.75.75 0 0 1-.947.373a4.375 4.375 0 1 0-1.586 8.455h11.648a2.978 2.978 0 1 0-.77-5.855a.75.75 0 0 1-.939-.812a3.91 3.91 0 0 0-3.882-4.374Zm-4.552.986a5.41 5.41 0 0 1 9.952 2.605a4.478 4.478 0 1 1 .191 8.951H6.875A5.875 5.875 0 1 1 8.38 7.194Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}