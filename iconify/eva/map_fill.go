package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMapFill0"><g id="evaMapFill1"><path id="evaMapFill2" fill="currentColor" d="m20.41 5.89l-4-1.8h-.82L12 5.7L8.41 4.09h-.05L8.24 4h-.6l-4 1.8a1 1 0 0 0-.64 1V19a1 1 0 0 0 .46.84A1 1 0 0 0 4 20a1 1 0 0 0 .41-.09L8 18.3l3.59 1.61h.05a.85.85 0 0 0 .72 0h.05L16 18.3l3.59 1.61A1 1 0 0 0 20 20a1 1 0 0 0 .54-.16A1 1 0 0 0 21 19V6.8a1 1 0 0 0-.59-.91ZM9 6.55l2 .89v10l-2-.89Zm10 10.9l-2-.89v-10l2 .89Z"/></g></g>`),
		g.Group(children),
	)
}