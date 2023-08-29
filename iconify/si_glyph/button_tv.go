package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m14.338 3.045l-5.33 1.002l-5.361-1.002c-1.452 0-2.631 1.328-2.631 2.966v5.023C1.016 12.672 2.195 14 3.647 14l5.361-1.031L14.338 14c1.452 0 2.631-1.328 2.631-2.966V6.011c0-1.638-1.179-2.966-2.631-2.966zM8.024 7.016H6.026v4.031H4.964V7.016h-1.98V6h5.04v1.016zm5.062 4.017h-1.127L9.962 5.965h1.3l1.268 3.666l1.231-3.666h1.294l-1.969 5.068z"/>`),
		g.Group(children),
	)
}