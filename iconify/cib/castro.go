package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0C7.161 0 0 7.161 0 16s7.161 16 16 16s16-7.161 16-16S24.839 0 16 0zm-.005 18.656c-3.645 0-3.645-5.473 0-5.473c3.651 0 3.651 5.473 0 5.473zm6.661 6.469l-2.683-3.719c5.303-3.876 2.553-12.267-4.009-12.256c-6.568.016-9.281 8.417-3.964 12.271l-2.688 3.724a11.384 11.384 0 0 1-4.161-12.719a11.372 11.372 0 0 1 10.823-7.864c6.277 0 11.365 5.088 11.365 11.364a11.338 11.338 0 0 1-4.683 9.199z"/>`),
		g.Group(children),
	)
}