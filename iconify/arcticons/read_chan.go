package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadChan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.442 8.488L15.34 7.132l-4.76 4.863v2.695l13.459 9.248L28.315 8.18l-1.583-1.94l-6.922-.74l-1.368 2.987Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.693 18.353l1.384-3.09l-4.818-4.806l-2.694-.025l-9.527 13.507l15.868 4.29l1.954-1.565l.806-6.914l-2.973-1.397Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.508 39.533l3.083 1.399l4.827-4.797l.038-2.694l-13.418-9.503l-4.408 15.764l1.556 1.96l6.911.838l1.41-2.967Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.31 29.469L6.928 32.56l4.823 4.8l2.694.024l9.593-13.446l-15.953-4.345l-1.952 1.566l-.8 6.916l2.976 1.394Z"/>`),
		g.Group(children),
	)
}