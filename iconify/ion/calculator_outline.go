package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="288" height="416" x="112" y="48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" rx="32" ry="32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M160.01 112H352v64H160.01z"/><circle cx="168" cy="248" r="24" fill="currentColor"/><circle cx="256" cy="248" r="24" fill="currentColor"/><circle cx="344" cy="248" r="24" fill="currentColor"/><circle cx="168" cy="328" r="24" fill="currentColor"/><circle cx="256" cy="328" r="24" fill="currentColor"/><circle cx="168" cy="408" r="24" fill="currentColor"/><circle cx="256" cy="408" r="24" fill="currentColor"/><rect width="48" height="128" x="320" y="304" fill="currentColor" rx="24" ry="24"/>`),
		g.Group(children),
	)
}