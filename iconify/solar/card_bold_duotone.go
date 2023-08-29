package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 20h4c3.771 0 5.657 0 6.828-1.172C22 17.657 22 15.771 22 12c0-.442-.002-1.608-.004-2H2c-.002.392 0 1.558 0 2c0 3.771 0 5.657 1.171 6.828C4.343 20 6.23 20 10 20Z" opacity=".5"/><path d="M9.995 4h4.01c3.781 0 5.672 0 6.846 1.116c.846.803 1.083 1.96 1.149 3.884v1H2V9c.066-1.925.303-3.08 1.149-3.884C4.323 4 6.214 4 9.995 4ZM12.5 15.25a.75.75 0 0 0 0 1.5H14a.75.75 0 0 0 0-1.5h-1.5Zm-6.5 0a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5H6Z"/></g>`),
		g.Group(children),
	)
}