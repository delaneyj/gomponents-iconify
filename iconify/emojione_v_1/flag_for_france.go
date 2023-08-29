package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFrance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M22 10h20v44H22z"/><path fill="#1b75bb" d="M10 10C3.373 10 0 14.925 0 21v22c0 6.075 3.373 11 10 11h12V10H10z"/><path fill="#ec1c24" d="M52 10H42v44h12c6.627 0 10-4.925 10-11V21c0-6.076-.042-11-12-11"/>`),
		g.Group(children),
	)
}