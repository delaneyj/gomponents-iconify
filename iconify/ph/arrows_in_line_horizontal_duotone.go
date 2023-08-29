package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInLineHorizontalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m64 96l32 32l-32 32Zm96 32l32 32V96Z" opacity=".2"/><path d="M136 40v176a8 8 0 0 1-16 0V40a8 8 0 0 1 16 0Zm-34.34 82.34a8 8 0 0 1 0 11.32l-32 32A8 8 0 0 1 56 160v-24H16a8 8 0 0 1 0-16h40V96a8 8 0 0 1 13.66-5.66Zm-17 5.66L72 115.31v25.38ZM248 128a8 8 0 0 1-8 8h-40v24a8 8 0 0 1-13.66 5.66l-32-32a8 8 0 0 1 0-11.32l32-32A8 8 0 0 1 200 96v24h40a8 8 0 0 1 8 8Zm-64-12.69L171.31 128L184 140.69Z"/></g>`),
		g.Group(children),
	)
}