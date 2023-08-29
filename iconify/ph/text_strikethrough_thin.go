package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextStrikethroughThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 128a4 4 0 0 1-4 4h-54.84c15.45 7.24 26.84 17.8 26.84 36c0 24.26-26.92 44-60 44s-60-19.74-60-44a4 4 0 0 1 8 0c0 19.85 23.33 36 52 36s52-16.15 52-36c0-19.54-16.13-28.3-42.18-36H40a4 4 0 0 1 0-8h176a4 4 0 0 1 4 4ZM76.33 100a3.85 3.85 0 0 0 1.25-.2a4 4 0 0 0 2.55-5a21.9 21.9 0 0 1-1-6.75c0-20.52 21-36 48.89-36c20.32 0 37 8.2 44.49 21.92a4 4 0 0 0 7-3.85C170.54 53.75 151.29 44 128 44c-32.43 0-56.89 18.92-56.89 44a29.76 29.76 0 0 0 1.42 9.25a4 4 0 0 0 3.8 2.75Z"/>`),
		g.Group(children),
	)
}