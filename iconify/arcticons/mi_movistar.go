package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiMovistar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.637 18.574C7.825 1.672.005 28.354 7.408 37.44c1.992 2.444 4.476-.096 3.797-2.491c-3.3-11.65-.335-11.838 3.43-4.696c3.094 5.87 11.94 13.457 16.21-1.143c2.967-10.148 7.58-10.052 7.43-3.92c-.255 10.42 2.598 8.595 4.492 3.349c3.456-9.579-6.007-28.959-13.678-12.7c-2.935 6.22-4.41 12.181-10.452 2.735Z"/>`),
		g.Group(children),
	)
}