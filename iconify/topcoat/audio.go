package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M27.5 9.045a13.391 13.391 0 0 1 8.37 12.425a13.39 13.39 0 0 1-8.37 12.424v2.795a16.007 16.007 0 0 0 11-15.219c0-7.099-4.609-13.125-11-15.228v2.803zm0 6.599a7.644 7.644 0 0 1 2.68 5.827a7.64 7.64 0 0 1-2.68 5.826v2.844c2.99-1.731 5-4.966 5-8.67a10.01 10.01 0 0 0-5-8.67v2.843zm-23 13.835h8l9 11.015c1 1.44 3.34 1.331 4-.302V2.83c-.811-1.632-2.939-1.763-4-.382l-9 11.053h-8c-2.561 0-3 .461-3 2.964v10.012c0 2.442.5 3.002 3 3.002z"/>`),
		g.Group(children),
	)
}