package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nayateljoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M22.032 29.341V18.203l5.566 5.569l-5.566 5.569z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-miterlimit="6" d="M3.708 16.878C5.75 26.571 12.56 33.684 20.66 33.684c0 0 8.75.353 15.007-8.951"/><path fill="none" stroke="currentColor" stroke-miterlimit="6" d="M39.578 9.15c-9.578-2.588-19.02.257-22.695 7.483c0 0-4.655 7.12.4 16.629"/><path fill="none" stroke="currentColor" stroke-miterlimit="6" d="M35.647 24.763a13.783 13.783 0 0 0-1.802-3.59s-4.328-7.28-14.608-7.88"/><path fill="none" stroke="currentColor" stroke-miterlimit="6" d="M29.822 44.702c5.622-6.018 7.89-13.708 5.825-19.94"/>`),
		g.Group(children),
	)
}