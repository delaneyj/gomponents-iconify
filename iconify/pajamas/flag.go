package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 10.174v.012c-1.393.72-2.81.247-4.959-.585l-.17-.066C6.763 8.91 4.685 8.105 2.5 8.63V2.814c1.393-.72 2.81-.247 4.959.585l.17.066c1.607.624 3.685 1.43 5.871.904v5.805ZM8 11c-1.83-.708-3.659-1.416-5.5-.81v4.06a.75.75 0 0 1-1.5 0V2C3.35.2 5.675 1.1 8 2c1.83.708 3.659 1.416 5.5.81A5.068 5.068 0 0 0 15 2v9c-2.35 1.8-4.675.9-7 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}