package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SirenThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M124 16V8a4 4 0 0 1 8 0v8a4 4 0 0 1-8 0Zm76 28a4 4 0 0 0 2.83-1.17l8-8a4 4 0 1 0-5.66-5.66l-8 8A4 4 0 0 0 200 44ZM53.17 42.83a4 4 0 0 0 5.66-5.66l-8-8a4 4 0 0 0-5.66 5.66Zm83.49 33.22a4 4 0 0 0-1.32 7.9C156.24 87.45 172 106.39 172 128a4 4 0 0 0 8 0c0-25.47-18.63-47.8-43.34-51.95ZM228 176v24a12 12 0 0 1-12 12H40a12 12 0 0 1-12-12v-24a12 12 0 0 1 12-12h4v-36a84 84 0 0 1 84-84h.64c46 .34 83.36 38.47 83.36 85v35h4a12 12 0 0 1 12 12ZM52 164h152v-35c0-42.15-33.83-76.69-75.42-77A76 76 0 0 0 52 128Zm168 12a4 4 0 0 0-4-4H40a4 4 0 0 0-4 4v24a4 4 0 0 0 4 4h176a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}