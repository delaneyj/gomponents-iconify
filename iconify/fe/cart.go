package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCart1" fill="currentColor"><path id="feCart2" d="M8 16h8a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1V4H5a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1v2.001L8.073 5h9.854C19.072 5 20 5.895 20 7c0 .146-.017.291-.05.434l-1.151 5c-.21.915-1.052 1.566-2.024 1.566H8.073L8 13.999V16Zm-.5 6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm9 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM8 7v5h8.831L18 7H8Z"/></g></g>`),
		g.Group(children),
	)
}