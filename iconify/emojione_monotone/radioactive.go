package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radioactive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.431 2 2 15.432 2 32c0 16.568 13.431 30 30 30c16.57 0 30-13.432 30-30C62 15.432 48.57 2 32 2zm0 57C17.088 59 5 46.912 5 32S17.088 5 32 5s27 12.089 27 27s-12.088 27-27 27z"/><circle cx="32" cy="32" r="5" fill="currentColor"/><path fill="currentColor" d="m8 28.76l17.12 2.221a6.924 6.924 0 0 1 2.668-4.62L17.304 12.644C12.303 16.473 8.843 22.205 8 28.76zm24 10.056a6.923 6.923 0 0 1-2.668-.533l-6.638 15.938A24.12 24.12 0 0 0 32 56.078a24.14 24.14 0 0 0 9.309-1.857L34.67 38.283a6.933 6.933 0 0 1-2.67.533m6.881-7.835L56 28.76c-.842-6.555-4.303-12.287-9.303-16.116L36.213 26.361a6.928 6.928 0 0 1 2.668 4.62"/>`),
		g.Group(children),
	)
}