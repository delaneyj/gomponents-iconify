package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetGeneration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.299 6.366C39.327 14.645 31.973 24.119 24 24.072C11.57 24 8.583 13.122 2.885 19.94m10.094 22.524c9.273 3.036 18.807-6.538 23.56-6.247c1.458.089 2.958.48 4.4 1.025"/>`),
		g.Group(children),
	)
}