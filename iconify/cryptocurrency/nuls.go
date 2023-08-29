package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nuls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-1.597-12.64l-.018 4.492l-2.825-1.447l.065-9.454l4.744 6.223l4.274 2.387L22 20.443V8.779L16.028 6l.064 6.8l1.413 1.873l.129-6.196l2.668 1.331v9.88l-3.019-1.846l-5.465-7.164a.636.636 0 0 0-.923-.098l-.683.586a.591.591 0 0 0-.212.453v11.718L16 26v-4.536z"/>`),
		g.Group(children),
	)
}