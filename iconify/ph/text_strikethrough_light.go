package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextStrikethroughLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 128a6 6 0 0 1-6 6h-46.55c11.28 6.92 20.55 17.38 20.55 34c0 25.36-27.81 46-62 46s-62-20.64-62-46a6 6 0 0 1 12 0c0 18.75 22.43 34 50 34s50-15.25 50-34c0-18.23-15.46-26.59-40.47-34H40a6 6 0 0 1 0-12h176a6 6 0 0 1 6 6ZM76.33 102a6.2 6.2 0 0 0 1.88-.3A6 6 0 0 0 82 94.13a19.74 19.74 0 0 1-.89-6.13c0-19.38 20.16-34 46.89-34c19.58 0 35.56 7.81 42.74 20.89a6 6 0 0 0 10.52-5.78C171.94 52.13 152 42 128 42c-33.57 0-58.89 19.77-58.89 46a31.62 31.62 0 0 0 1.52 9.87a6 6 0 0 0 5.7 4.13Z"/>`),
		g.Group(children),
	)
}