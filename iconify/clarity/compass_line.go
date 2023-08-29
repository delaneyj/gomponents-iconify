package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M20.82 15.31L10.46 9c-.46-.26-1.11.37-.86.84l6.15 10.56l10.56 6.15a.66.66 0 0 0 .84-.86Zm-4 4l3-3l4.55 7.44Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm1 29.95v-2.42h-2v2.42A14 14 0 0 1 4.05 19h2.42v-2H4.05A14 14 0 0 1 17 4.05v2.42h2V4.05A14 14 0 0 1 31.95 17h-2.42v2h2.42A14 14 0 0 1 19 31.95Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}