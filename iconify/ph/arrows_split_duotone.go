package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsSplitDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M32 184h64l-32 32Zm128 0l32 32l32-32Z" opacity=".2"/><path d="M231.39 180.94A8 8 0 0 0 224 176h-24v-40a8 8 0 0 0-2.34-5.66L136 68.69V24a8 8 0 0 0-16 0v44.69l-61.66 61.65A8 8 0 0 0 56 136v40H32a8 8 0 0 0-5.66 13.66l32 32a8 8 0 0 0 11.32 0l32-32A8 8 0 0 0 96 176H72v-36.69l56-56l56 56V176h-24a8 8 0 0 0-5.66 13.66l32 32a8 8 0 0 0 11.32 0l32-32a8 8 0 0 0 1.73-8.72ZM64 204.69L51.31 192h25.38Zm128 0L179.31 192h25.38Z"/></g>`),
		g.Group(children),
	)
}