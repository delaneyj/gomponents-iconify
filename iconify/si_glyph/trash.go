package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m1.982 5.073l1.025 10.266c0 .366.307.661.684.661h7.58a.673.673 0 0 0 .684-.661L12.98 5.073H1.982zm6.051 8.995H6.961V6.989h1.072v7.079zm2 0H8.961l1-7.079h1.072l-1 7.079zm-4 0H4.961l-1-7.079h1.072l1 7.079zm7.042-11.963H9.937V.709C9.937.317 9.481 0 9.081 0H5.986c-.4 0-.955.225-.955.615v1.396l-3.145.094a.717.717 0 0 0-.727.708v1.155H13.8V2.813a.715.715 0 0 0-.725-.708zM5.947 1.44c0-.312.351-.565.783-.565h1.564c.432 0 .782.254.782.565v.665h-3.13V1.44h.001z"/>`),
		g.Group(children),
	)
}