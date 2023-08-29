package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLinesDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m224 136l-96 96l-96-96h48v-32h96v32Z" opacity=".2"/><path d="M231.39 132.94A8 8 0 0 0 224 128h-40v-24a8 8 0 0 0-8-8H80a8 8 0 0 0-8 8v24H32a8 8 0 0 0-5.66 13.66l96 96a8 8 0 0 0 11.32 0l96-96a8 8 0 0 0 1.73-8.72ZM128 220.69L51.31 144H80a8 8 0 0 0 8-8v-24h80v24a8 8 0 0 0 8 8h28.69ZM72 40a8 8 0 0 1 8-8h96a8 8 0 0 1 0 16H80a8 8 0 0 1-8-8Zm0 32a8 8 0 0 1 8-8h96a8 8 0 0 1 0 16H80a8 8 0 0 1-8-8Z"/></g>`),
		g.Group(children),
	)
}