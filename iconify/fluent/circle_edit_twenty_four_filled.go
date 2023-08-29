package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleEditTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M20.715 11c.427 0 .855.084 1.257.25A10 10 0 0 0 12 2C6.477 2 2 6.477 2 12c0 5.184 3.945 9.447 8.997 9.95a2.117 2.117 0 0 1 .064-.543l.458-1.83c.162-.648.497-1.24.97-1.712l5.902-5.903A3.28 3.28 0 0 1 20.713 11h.002Z"/><path d="M20.715 12h-.002c-.585 0-1.17.223-1.615.67l-5.902 5.902a2.684 2.684 0 0 0-.707 1.247l-.458 1.831a1.087 1.087 0 0 0 1.319 1.318l1.83-.457a2.684 2.684 0 0 0 1.248-.707l5.902-5.902A2.285 2.285 0 0 0 20.715 12Z"/></g>`),
		g.Group(children),
	)
}