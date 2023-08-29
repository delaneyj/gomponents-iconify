package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInCardinalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 96L96 64h64Zm-32 96h64l-32-32Zm64-64l32 32V96ZM64 96v64l32-32Z" opacity=".2"/><path d="M122.34 101.66a8 8 0 0 0 11.32 0l32-32A8 8 0 0 0 160 56h-24V24a8 8 0 0 0-16 0v32H96a8 8 0 0 0-5.66 13.66ZM140.69 72L128 84.69L115.31 72Zm-7 82.34a8 8 0 0 0-11.32 0l-32 32A8 8 0 0 0 96 200h24v32a8 8 0 0 0 16 0v-32h24a8 8 0 0 0 5.66-13.66ZM115.31 184L128 171.31L140.69 184ZM232 120h-32V96a8 8 0 0 0-13.66-5.66l-32 32a8 8 0 0 0 0 11.32l32 32A8 8 0 0 0 200 160v-24h32a8 8 0 0 0 0-16Zm-48 20.69L171.31 128L184 115.31Zm-82.34-18.35l-32-32A8 8 0 0 0 56 96v24H24a8 8 0 0 0 0 16h32v24a8 8 0 0 0 13.66 5.66l32-32a8 8 0 0 0 0-11.32ZM72 140.69v-25.38L84.69 128Z"/></g>`),
		g.Group(children),
	)
}