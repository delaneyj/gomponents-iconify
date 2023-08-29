package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileShortcut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 23H4V1h16v5h-2V3H6v18h12v-3h2v5Zm-6.996-3.996H11V17h2.004v2.004Zm8.883-2.717l1.568-.781l-1.568-.78l-.781-1.57l-.781 1.57l-1.57.78l1.57.78l.78 1.57l.782-1.57Zm-5.83-.986l-1.097-2.204L12.756 12l2.204-1.098l1.097-2.204l1.097 2.204L19.358 12l-2.204 1.097l-1.097 2.204Zm5.83-6.024l1.568-.78l-1.568-.782l-.781-1.568l-.781 1.568l-1.57.781l1.57.781l.78 1.569l.782-1.569Z"/>`),
		g.Group(children),
	)
}