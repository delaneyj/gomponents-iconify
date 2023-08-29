package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SbbInclusive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.17 27.561c-2.116 1.063-3.63 3.916-3.63 7.268h10.92c0-3.352-1.513-6.204-3.628-7.267"/><ellipse cx="24" cy="25.666" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.273" ry="3.21"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.333 30.623C9.856 31.38 9 32.278 9 33.24c0 2.697 6.716 4.883 15 4.883s15-2.186 15-4.883c0-.963-.856-1.86-2.333-2.616m-14.922-5.356a4.834 4.834 0 0 0-.75-.47"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.332 24.797c-2.116 1.063-3.63 3.915-3.63 7.267H18.9"/><ellipse cx="19.163" cy="22.901" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.273" ry="3.21"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.099 32.064h5.199c0-3.351-1.514-6.203-3.629-7.267m-3.662 0a4.837 4.837 0 0 0-.752.471"/><ellipse cx="28.837" cy="22.901" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.273" ry="3.21"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.87 6.041l2.865 2.868l-2.867 2.867m-5.739 0L18.265 8.91l2.867-2.868m-2.867 2.867h11.47M24 6.041v5.736"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.421 15.318h39.158"/>`),
		g.Group(children),
	)
}