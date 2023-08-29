package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PricetagsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPricetagsFill0"><g id="evaPricetagsFill1"><path id="evaPricetagsFill2" fill="currentColor" d="m21.47 11.58l-6.42-6.41a1 1 0 0 0-.61-.29L5.09 4a1 1 0 0 0-.8.29a1 1 0 0 0-.29.8l.88 9.35a1 1 0 0 0 .29.61l6.41 6.42a1.84 1.84 0 0 0 1.29.53a1.82 1.82 0 0 0 1.28-.53l7.32-7.32a1.82 1.82 0 0 0 0-2.57Zm-9.91 0a1.5 1.5 0 1 1 0-2.12a1.49 1.49 0 0 1 0 2.1Z"/></g></g>`),
		g.Group(children),
	)
}