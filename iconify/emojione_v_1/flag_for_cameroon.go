package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCameroon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M22 10h20v44H22z"/><path fill="#007a5e" d="M10 10C3.373 10 0 14.925 0 21v22c0 6.075 3.373 11 10 11h12V10H10z"/><path fill="#f9cb38" d="M54 10H42v44h12c6.627 0 10-4.925 10-11V21c0-6.075-3.373-11-10-11M40.462 28.689l-6.202.011l-1.92-6.274l-1.914 6.274l-6.206-.011l5.03 3.822l-1.952 6.232l5.05-3.871l5.05 3.871l-1.958-6.232z"/>`),
		g.Group(children),
	)
}