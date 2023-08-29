package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Corkscrew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.107.038H5.859C5.386.038 5 .448 5 .953v.085c0 .504.387.915.859.915h7.248c.476 0 .859-.411.859-.915V.953c.001-.505-.383-.915-.859-.915zm-2.12 8.458c.012-.146-1.021-1.309-1.021-1.309l-.025-4.232h2.012V2h-4.91v.955h2.014v1.528s-1.015.886-1.017 1.024c-.005.137.981 1.326.981 1.326v4.338s-.951 1.215-.953 1.374c0 .159.957 1.329.957 1.329l.012 1.625c0 .277.215.5.48.5c.268 0 .451-.191.451-.469l-.002-5.636c.001.002 1.01-1.251 1.021-1.398z"/>`),
		g.Group(children),
	)
}