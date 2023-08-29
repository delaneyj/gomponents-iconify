package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerAltPoi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M31.861 31.44L1.383 48.165a4.725 2.593 0 0 0 0 3.668L46.658 76.68a4.725 2.593 0 0 0 6.684 0l45.275-24.846a4.725 2.593 0 0 0 0-3.668l-30.451-16.71a1408.29 1408.29 0 0 1-3.236 5.558L88.596 50L50 71.18L11.406 50l23.698-13.006a625.903 625.903 0 0 0-3.243-5.555z" color="currentColor"/><path fill="currentColor" d="M50.001 0c-10.09 0-18.359 8.245-18.359 18.306a18.13 18.13 0 0 0 3.353 10.5l12.766 22.068c1.788 2.336 2.976 1.89 4.462-.124l14.08-23.96c.284-.514.508-1.063.702-1.623a18.062 18.062 0 0 0 1.353-6.86C68.358 8.244 60.092 0 50 0Zm0 8.577c5.434 0 9.754 4.311 9.754 9.73c0 5.417-4.32 9.726-9.754 9.726c-5.433 0-9.756-4.309-9.756-9.727S44.568 8.577 50 8.577z" color="currentColor"/>`),
		g.Group(children),
	)
}