package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunDimBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M116 36v-4a12 12 0 0 1 24 0v4a12 12 0 0 1-24 0Zm80 92a68 68 0 1 1-68-68a68.07 68.07 0 0 1 68 68Zm-24 0a44 44 0 1 0-44 44a44.05 44.05 0 0 0 44-44ZM51.51 68.49a12 12 0 1 0 17-17l-4-4a12 12 0 0 0-17 17Zm0 119l-4 4a12 12 0 0 0 17 17l4-4a12 12 0 1 0-17-17ZM196 72a12 12 0 0 0 8.49-3.51l4-4a12 12 0 0 0-17-17l-4 4A12 12 0 0 0 196 72Zm8.49 115.51a12 12 0 0 0-17 17l4 4a12 12 0 0 0 17-17ZM48 128a12 12 0 0 0-12-12h-4a12 12 0 0 0 0 24h4a12 12 0 0 0 12-12Zm80 80a12 12 0 0 0-12 12v4a12 12 0 0 0 24 0v-4a12 12 0 0 0-12-12Zm96-92h-4a12 12 0 0 0 0 24h4a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}