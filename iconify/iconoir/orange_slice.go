package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeSlice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.613 10.11l7.778-7.777c4.295 4.296 4.295 11.26 0 15.556c-4.296 4.296-11.261 4.296-15.557 0l7.779-7.778Zm0 0l-.354 8.133m.354-8.132h7.778m-7.778 0l5.303 5.303"/>`),
		g.Group(children),
	)
}