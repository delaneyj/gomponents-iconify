package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsVerticalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M160 56H96l32-32Zm-32 176l32-32H96Z" opacity=".2"/><path d="M160 192h-24V64h24a8 8 0 0 0 5.66-13.66l-32-32a8 8 0 0 0-11.32 0l-32 32A8 8 0 0 0 96 64h24v128H96a8 8 0 0 0-5.66 13.66l32 32a8 8 0 0 0 11.32 0l32-32A8 8 0 0 0 160 192ZM128 35.31L140.69 48h-25.38Zm0 185.38L115.31 208h25.38Z"/></g>`),
		g.Group(children),
	)
}