package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.001 1.734a1 1 0 0 1 .501.135l16.004 9.265a1 1 0 0 1 0 1.731L4.502 22.131a1 1 0 0 1-1.501-.866V2.735a1 1 0 0 1 1-1Zm8.292 11.68l-4.498 4.498l5.699-3.299l-1.2-1.2ZM5 6.118v11.76l5.88-5.88L5 6.12Zm10.284 4.302L13.707 12l1.578 1.577L18.009 12l-2.725-1.579Zm-7.49-4.336l4.5 4.5l1.199-1.2l-5.699-3.3Z"/>`),
		g.Group(children),
	)
}