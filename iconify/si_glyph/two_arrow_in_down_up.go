package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoArrowInDownUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.771 4.096a.664.664 0 0 0 0-.946a.678.678 0 0 0-.954 0L9.989 4.503V1c0-.553-.444-1-.992-1a.995.995 0 0 0-.992 1v3.564L6.187 3.18a.674.674 0 0 0-1.15.473c0 .171.066.343.199.474l3.74 2.819l3.795-2.85zm0 7.78a.664.664 0 0 1 0 .946a.678.678 0 0 1-.954 0l-1.828-1.328V15c0 .553-.444 1-.992 1a.995.995 0 0 1-.992-1v-3.428l-1.818 1.266a.674.674 0 0 1-1.15-.473c0-.171.066-.343.199-.474l3.74-2.838l3.795 2.823z"/>`),
		g.Group(children),
	)
}