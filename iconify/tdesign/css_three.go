package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.941 1H22.06l-1.098 19.208l-8.96 3.36l-8.962-3.36L1.941 1ZM4.06 3l.902 15.792l7.04 2.64l7.038-2.64L19.941 3H4.06Zm1.423 2H18.33l-.108 1.887l-7.29 3.49h7.09l-.404 7.084L12 19.568l-5.618-2.107l-.114-1.994v-.004l-.075-1.39h2.01l.062 1.284l.04.69L12 17.431l3.696-1.386l.21-3.67H6.09l-.106-1.848L13.355 7H12l-6.255.012L5.482 5Z"/>`),
		g.Group(children),
	)
}