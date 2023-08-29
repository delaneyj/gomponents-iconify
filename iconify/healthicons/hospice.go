package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M41.708 23.794L24.058 6.086L6.294 23.792a1 1 0 1 0 1.412 1.416L11 21.925V41a1 1 0 0 0 1 1h24a1 1 0 0 0 1-1V21.904l3.292 3.302a1 1 0 1 0 1.416-1.412ZM20.25 20C17.794 20 16 22.655 16 25.517c0 .77.116 1.506.32 2.205c.1.342.22.676.36 1C18.61 33.244 24 36 24 36s8-4.355 8-10.483C32 22.655 30.206 20 27.75 20c-1.705 0-2.97 1.191-3.75 2.88c-.78-1.689-2.046-2.88-3.75-2.88Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}