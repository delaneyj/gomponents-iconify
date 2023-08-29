package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestaurantTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.268 12.147l-.853.853l7.07 7.071l-1.413 1.414L12 14.415l-7.072 7.07l-1.414-1.414l9.339-9.339c-.588-1.457.02-3.555 1.621-5.156c1.953-1.953 4.644-2.428 6.01-1.061c1.368 1.367.893 4.058-1.06 6.01c-1.602 1.602-3.7 2.21-5.157 1.622ZM4.222 3.807l6.718 6.718l-2.829 2.829l-3.889-3.89a4 4 0 0 1 0-5.656Zm13.789 5.304c1.257-1.257 1.516-2.726 1.06-3.182c-.455-.456-1.924-.196-3.181 1.06c-1.258 1.258-1.517 2.727-1.061 3.183c.456.455 1.925.196 3.182-1.06Z"/>`),
		g.Group(children),
	)
}