package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutCardinalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m128 24l32 32H96Zm0 208l32-32H96Zm72-136v64l32-32ZM24 128l32 32V96Z" opacity=".2"/><path d="M96 64h24v32a8 8 0 0 0 16 0V64h24a8 8 0 0 0 5.66-13.66l-32-32a8 8 0 0 0-11.32 0l-32 32A8 8 0 0 0 96 64Zm32-28.69L140.69 48h-25.38ZM160 192h-24v-32a8 8 0 0 0-16 0v32H96a8 8 0 0 0-5.66 13.66l32 32a8 8 0 0 0 11.32 0l32-32A8 8 0 0 0 160 192Zm-32 28.69L115.31 208h25.38Zm109.66-98.35l-32-32A8 8 0 0 0 192 96v24h-32a8 8 0 0 0 0 16h32v24a8 8 0 0 0 13.66 5.66l32-32a8 8 0 0 0 0-11.32ZM208 140.69v-25.38L220.69 128ZM96 136a8 8 0 0 0 0-16H64V96a8 8 0 0 0-13.66-5.66l-32 32a8 8 0 0 0 0 11.32l32 32A8 8 0 0 0 64 160v-24Zm-48 4.69L35.31 128L48 115.31Z"/></g>`),
		g.Group(children),
	)
}