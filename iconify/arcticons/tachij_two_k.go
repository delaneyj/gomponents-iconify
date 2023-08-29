package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TachijTwoK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.757 34.53h26.486M12.139 14.729h23.722M27.243 34.53c4.68-6.901 6.357-9.378 5.581-16.68m-17.842 0c-.856 6.605 3.044 12.531 4.345 14.603m6.779-17.724l1.607-3.284l-2.472.721L24 9.22l-1.241 2.946l-2.472-.721l1.607 3.284"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}