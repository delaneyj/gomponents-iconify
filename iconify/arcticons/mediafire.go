package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mediafire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.39 13.54c6 .2 12.18 3.71 12.11 11c-.06 6.2-4.72 10-12.1 10c-5.71 0-10.09-4.72-15.92-5.38c-4.21-.47 3.39-.2 1.89-2.55c-1.36-2.13-5.63-2.49-8.2-2.1c-3.74.57-4.41 2-4.67 2.6c2.51-5.8 6.13-3.67 8-6.58c.74-1.11-1.72-1.57-3.37-.35a7.06 7.06 0 0 0-.67.67a3.86 3.86 0 0 1 .67-.67a8.89 8.89 0 0 1 5.94-2.22c5.85-.27 11.19 2.93 11.73.28s-5-1.32-4.22-2a12.6 12.6 0 0 1 7.57-2.63Zm2.23 7c-2.37.2-2.45 1.24-4.69 2.44c-3.88 2.06-6.19 1.17-6.19 1.28s1.58.74 5.42 2.61a13.82 13.82 0 0 0 5.46 1.59c2.82.08 4.81-1.71 4.81-3.9a4.25 4.25 0 0 0-4.81-4.05Z"/>`),
		g.Group(children),
	)
}