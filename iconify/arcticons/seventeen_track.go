package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeventeenTrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.602 24.617l-6.735 12.516h6.735ZM8.219 14.443l6.566-3.576m0 0v26.266"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".998" d="m26.326 37.133l14.133-26.266H23.058"/>`),
		g.Group(children),
	)
}