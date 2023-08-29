package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PricetagsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPricetagsOutline0"><g id="evaPricetagsOutline1"><g id="evaPricetagsOutline2" fill="currentColor"><path d="M12.87 22a1.84 1.84 0 0 1-1.29-.53l-6.41-6.42a1 1 0 0 1-.29-.61L4 5.09a1 1 0 0 1 .29-.8a1 1 0 0 1 .8-.29l9.35.88a1 1 0 0 1 .61.29l6.42 6.41a1.82 1.82 0 0 1 0 2.57l-7.32 7.32a1.82 1.82 0 0 1-1.28.53Zm-6-8.11l6 6l7.05-7.05l-6-6l-7.81-.73Z"/><circle cx="10.5" cy="10.5" r="1.5"/></g></g></g>`),
		g.Group(children),
	)
}