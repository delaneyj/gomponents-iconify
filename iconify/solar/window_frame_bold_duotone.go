package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowFrameBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 2c4.714 0 7.071 0 8.535 1.464c1.08 1.08 1.364 2.647 1.439 5.286L22 9.5H2.026v-.75c.075-2.64.358-4.205 1.438-5.286C4.93 2 7.286 2 12 2Z" opacity=".5"/><path d="M13 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M2 12c0 4.714 0 7.071 1.464 8.535c1.01 1.01 2.446 1.324 4.786 1.421L9 22V9.5H2.026l-.023.75C2 10.794 2 11.377 2 12Z" opacity=".7"/><path d="M22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22c-.819 0-2.316 0-3-.008V9.5h13l-.003.75C22 10.794 22 11.377 22 12Z"/></g>`),
		g.Group(children),
	)
}