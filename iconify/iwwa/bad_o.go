package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M20.001 2.683C10.452 2.683 2.684 10.451 2.684 20s7.769 17.317 17.317 17.317S37.316 29.548 37.316 20S29.549 2.683 20.001 2.683zm0 33.134c-8.722 0-15.817-7.096-15.817-15.817S11.279 4.183 20.001 4.183c8.721 0 15.815 7.096 15.815 15.817s-7.094 15.817-15.815 15.817z"/><path fill="currentColor" d="M20.013 22.697a7.726 7.726 0 0 0-6.604 3.682a.375.375 0 0 0 .122.516l.572.353a.375.375 0 0 0 .515-.122a6.304 6.304 0 0 1 5.394-3.011a6.305 6.305 0 0 1 5.374 2.979a.376.376 0 0 0 .516.118l.568-.356a.372.372 0 0 0 .118-.516a7.72 7.72 0 0 0-6.575-3.643z"/><circle cx="15.431" cy="16.005" r="1.044" fill="currentColor"/><circle cx="24.568" cy="16.005" r="1.044" fill="currentColor"/>`),
		g.Group(children),
	)
}