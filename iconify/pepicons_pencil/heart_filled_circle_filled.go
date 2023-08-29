package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartFilledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilHeartFilledCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 7.543c-1.25-.98-2.965-1.245-4.432-.895C6.678 7.1 5 8.621 5 11.163c0 3.326 2.88 6.016 7.571 8.24a1 1 0 0 0 .857 0C18.12 17.18 21 14.49 21 11.164c0-2.542-1.678-4.064-3.568-4.515c-1.467-.35-3.183-.084-4.432.895Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilHeartFilledCircleFilled0)"/></g>`),
		g.Group(children),
	)
}