package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayStreamLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M19.141 5A9.967 9.967 0 0 1 22 12a9.969 9.969 0 0 1-2.922 7.064M5 19.142A9.97 9.97 0 0 1 2 12a9.969 9.969 0 0 1 2.936-7.078m11.349 3.122C17.345 9.059 18 10.449 18 11.982c0 1.552-.67 2.957-1.753 3.974M7.8 16C6.69 14.979 6 13.556 6 11.982C6 10.427 6.673 9.018 7.762 8"/><path d="M13.656 10.451C14.552 11.11 15 11.438 15 12c0 .562-.448.891-1.344 1.549a13.08 13.08 0 0 1-.718.495a12.23 12.23 0 0 1-.653.38c-.894.49-1.34.735-1.741.464c-.401-.271-.438-.838-.51-1.971c-.02-.321-.034-.635-.034-.917c0-.282.013-.596.034-.917c.072-1.133.109-1.7.51-1.97c.4-.272.847-.027 1.74.462c.233.127.457.256.654.381c.226.143.471.314.718.495Z"/></g>`),
		g.Group(children),
	)
}