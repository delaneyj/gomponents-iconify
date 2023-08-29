package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassportOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 9.25a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5ZM9.75 13a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z" clip-rule="evenodd"/><path d="M10 18.25a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5h-4Z"/><path fill-rule="evenodd" d="M15.611 1.584a2.75 2.75 0 0 1 3.139 2.722v1.377a3.75 3.75 0 0 1 2 3.317v10A3.75 3.75 0 0 1 17 22.75H7A3.75 3.75 0 0 1 3.25 19V5h.001l-.001-.071c0-.906.666-1.675 1.564-1.803L15.61 1.584ZM5.071 6.75H4.75V19A2.25 2.25 0 0 0 7 21.25h10A2.25 2.25 0 0 0 19.25 19V9a2.25 2.25 0 0 0-2.24-2.25H5.07Zm12.179-1.5H5.067a.321.321 0 0 1-.041-.639L15.823 3.07a1.25 1.25 0 0 1 1.427 1.237v.944Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}