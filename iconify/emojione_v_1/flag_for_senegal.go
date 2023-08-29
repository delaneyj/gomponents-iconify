package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSenegal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9cb38" d="M18 10h28v44H18z"/><path fill="#137a08" d="M10 10C3.373 10 0 14.925 0 21v22c0 6.075 3.373 11 10 11h10V10H10z"/><path fill="#ec1c24" d="M54 10H44v44h10c6.627 0 10-4.925 10-11V21c0-6.075-3.373-11-10-11"/><path fill="#137a08" d="m41.11 28.22l-6.679.01l-2.074-6.471L30.3 28.23l-6.695-.01l5.425 3.95l-2.098 6.433l5.442-4l5.437 4l-2.101-6.433z"/>`),
		g.Group(children),
	)
}