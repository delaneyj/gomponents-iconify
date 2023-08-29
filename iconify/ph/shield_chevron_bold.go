package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldChevronBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H48a20 20 0 0 0-20 20v58.8c0 92.36 78.1 123 93.76 128.18a19.6 19.6 0 0 0 12.48 0C149.9 237.78 228 207.16 228 114.8V56a20 20 0 0 0-20-20ZM52 60h152v54.8a113.07 113.07 0 0 1-6.6 39.14l-62.52-43.77a12 12 0 0 0-13.76 0L58.6 153.94A113.07 113.07 0 0 1 52 114.8Zm76 159.75c-10.07-3.53-39.26-15.79-58.39-44.22L128 134.65l58.39 40.88C167.26 204 138.07 216.22 128 219.75Z"/>`),
		g.Group(children),
	)
}