package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.5 14h-9c-.83 0-1.5-.67-1.5-1.5v-9C2 2.67 2.67 2 3.5 2h9c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5Zm-9-11c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h9c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5h-9Z"/>`),
		g.Group(children),
	)
}