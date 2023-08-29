package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M79.437 47.349c-3.51-12.982-14.464-22.083-27.134-22.083c-8.835 0-17.065 4.454-22.414 12.018a20.179 20.179 0 0 0-4.501-.514c-11.889 0-21.563 10.539-21.563 23.498c0 4.647 1.251 9.148 3.612 13.018c.555.906 1.49 1.449 2.49 1.449h80.84c.947 0 1.836-.485 2.403-1.315a17.249 17.249 0 0 0 3.004-9.767c.001-9.658-7.814-17.473-16.737-16.304z"/>`),
		g.Group(children),
	)
}