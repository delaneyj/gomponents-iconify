package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DartSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.207.13a.74.74 0 0 1 .517.55l1.199 5.397l5.396 1.2a.741.741 0 0 1 .364 1.247l-7.159 7.159a.741.741 0 0 1-.685.2L1.982 14.58a.741.741 0 0 1-.563-.563L.118 8.16a.741.741 0 0 1 .2-.685L7.475.317A.741.741 0 0 1 8.207.13Zm.283 6.332l-.92-4.141l-5.917 5.915l.92 4.142L8.49 6.462Zm-4.868 6.964l4.142.92l5.915-5.915l-4.141-.92l-5.916 5.915Z"/>`),
		g.Group(children),
	)
}