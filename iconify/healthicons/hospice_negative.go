package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospiceNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHospiceNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM23.058 6.086l17.65 17.708a1 1 0 0 1-1.416 1.412L37 22.906V41a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1V22.922l-2.294 2.286a1 1 0 1 1-1.412-1.416L23.058 6.086ZM35 20.9L23.053 8.914L11 20.928V40h24V20.9ZM19.25 20C16.794 20 15 22.655 15 25.517c0 .77.116 1.506.32 2.205c.1.342.22.676.36 1C17.61 33.244 23 36 23 36s8-4.355 8-10.483C31 22.655 29.206 20 26.75 20c-1.705 0-2.97 1.191-3.75 2.88c-.78-1.689-2.046-2.88-3.75-2.88Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHospiceNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}