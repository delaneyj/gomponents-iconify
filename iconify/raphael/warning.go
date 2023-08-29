package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.225 23.567l-3.778-6.542c-1.14-1.972-3.002-5.2-4.14-7.172l-3.78-6.542c-1.14-1.972-3.002-1.972-4.14 0L9.608 9.854l-4.143 7.172l-3.777 6.542c-1.14 1.974-.207 3.587 2.07 3.587H27.15c2.28 0 3.21-1.613 2.073-3.587zm-12.69 1.013h-2.24v-2.15h2.24v2.15zm-.107-3.736h-2.023l-.2-9.204h2.406l-.182 9.204z"/>`),
		g.Group(children),
	)
}