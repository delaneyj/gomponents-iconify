package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolarFlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.8 21.457h18.478c2.688 0 2.688-4.442 0-4.442H9.248m-.023-.008c2.725-5.47 8.38-8.936 14.576-8.936c8.965 0 16.232 7.144 16.232 15.955h0c0 8.812-7.267 15.956-16.232 15.956h0c-8.964 0-16.232-7.144-16.232-15.956h0c0-.861.071-1.722.213-2.572"/>`),
		g.Group(children),
	)
}