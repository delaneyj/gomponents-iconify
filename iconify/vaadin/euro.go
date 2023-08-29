package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.89 3a5.45 5.45 0 0 1 3.127 1.011L14 1.69a7.369 7.369 0 0 0-3.129-.686a7.482 7.482 0 0 0-7.323 5.947L2 7v1h1.41v.5a3.848 3.848 0 0 0 0 .512L1.999 9v1h1.54c.882 3.353 3.805 5.818 7.331 5.999a7.42 7.42 0 0 0 3.175-.708L14 13a5.426 5.426 0 0 1-3.108 1a5.909 5.909 0 0 1-5.28-3.959L12 10V9H5.41a3.224 3.224 0 0 1 .001-.511L5.41 8H12V7H5.6c.678-2.325 2.788-3.996 5.29-4z"/>`),
		g.Group(children),
	)
}