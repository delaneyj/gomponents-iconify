package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCompassFill0"><g id="evaCompassFill1"><g id="evaCompassFill2" fill="currentColor"><path d="m10.8 13.21l1.69-.68l.71-1.74l-1.69.68l-.71 1.74z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm3.93 7.42l-1.75 4.26a1 1 0 0 1-.55.55l-4.21 1.7A1 1 0 0 1 9 16a1 1 0 0 1-.71-.31h-.05a1 1 0 0 1-.18-1l1.75-4.26a1 1 0 0 1 .55-.55l4.21-1.7a1 1 0 0 1 1.1.25a1 1 0 0 1 .26.99Z"/></g></g></g>`),
		g.Group(children),
	)
}