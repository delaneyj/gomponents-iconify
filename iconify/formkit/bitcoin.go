package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.19 7.47c.76-.39 1.24-1.07 1.13-2.21c-.15-1.56-1.49-2.08-3.19-2.23V1H6.82v2H5.77V1H4.46v2c-.29 0-.56.01-.84.01H2v1.43s.77-.02.75 0c.53 0 .71.31.76.58v2.46h.14h-.14v3.44c-.02.17-.12.44-.49.44c.02.01-.75 0-.75 0l-.26 1.63h2.44v2h1.31v-2h1.05v2h1.31v-2.05c2.21-.13 3.75-.68 3.95-2.76c.16-1.67-.63-2.42-1.88-2.72Zm-4.4-2.94h.27c.92-.02 2.8-.07 2.8 1.32C8.86 7.21 6.9 7.18 6 7.16h-.22V4.54Zm3.68 5.39c0 1.5-2.35 1.46-3.43 1.45h-.26V8.48h.32c1.1-.03 3.36-.08 3.36 1.45Z"/>`),
		g.Group(children),
	)
}