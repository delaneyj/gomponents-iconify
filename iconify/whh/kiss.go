package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kiss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M608 568H416q-38 0-78-10.5T266.5 532t-71-41t-65-45.5t-67-52T0 344q38-29 95-84.5t98-97t94-74T384 56q39 0 72.5 17t55.5 47q22-30 55.5-47T640 56q44 0 97 32.5t94 74t98 97t95 84.5q-22 17-63.5 49.5t-67 52t-65 45.5t-71 41t-71.5 25.5t-78 10.5zm-32-320h-96q-23 0-46.5 9t-59 27t-63.5 29h9q1 26 57 44.5T512 376t135-18.5t57-44.5h24q-24-10-55-28t-52.5-27.5T576 248z"/>`),
		g.Group(children),
	)
}