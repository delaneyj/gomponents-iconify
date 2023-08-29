package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M592.09.004C442.66-.708 311.458 88.89 248.073 228.885c-32.601 77.731-30.662 184.188-29.15 277.147H52.297V1200h1095.406V506.032h-181.2V374.417c0-51.599-10.623-99.226-30.469-145.532C878.015 95.938 741.52.716 592.09.004zm0 220.971c87.291 2.317 150.961 67.954 154.76 152.124v132.933H438.575V374.417c8.239-91.624 66.223-155.759 153.515-153.442z"/>`),
		g.Group(children),
	)
}