package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 6.25a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5ZM10.25 9.5a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0Z"/><path d="M12 2.25a7.25 7.25 0 0 0-6.063 11.226l-2.587 4.48a.75.75 0 0 0 .795 1.11l2.614-.514l.861 2.52a.75.75 0 0 0 1.36.133l2.58-4.468a7.313 7.313 0 0 0 .88 0l2.58 4.468a.75.75 0 0 0 1.36-.134l.858-2.526l2.616.52a.75.75 0 0 0 .796-1.11l-2.586-4.479A7.25 7.25 0 0 0 12 2.25ZM6.25 9.5a5.75 5.75 0 1 1 11.5 0a5.75 5.75 0 0 1-11.5 0Zm3.734 6.966a7.243 7.243 0 0 1-3.027-1.757l-1.482 2.567l1.637-.322a.75.75 0 0 1 .854.493l.54 1.579l1.478-2.56Zm5.508 2.556l-1.476-2.556a7.244 7.244 0 0 0 3.027-1.757l1.48 2.563l-1.638-.326a.75.75 0 0 0-.856.495l-.537 1.58Z"/></g>`),
		g.Group(children),
	)
}