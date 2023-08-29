package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunglassesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 40a8 8 0 0 0 0 16a16 16 0 0 1 16 16v56H40V72a16 16 0 0 1 16-16a8 8 0 0 0 0-16a32 32 0 0 0-32 32v92a44 44 0 0 0 88 0v-20h32v20a44 44 0 0 0 88 0V72a32 32 0 0 0-32-32ZM93.66 181.66a8 8 0 0 1-11.32 0l-24-24a8 8 0 0 1 11.32-11.32l24 24a8 8 0 0 1 0 11.32Zm120 0a8 8 0 0 1-11.32 0l-24-24a8 8 0 0 1 11.32-11.32l24 24a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}