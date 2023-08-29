package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a30.16 30.16 0 0 1 5.25.44a23.87 23.87 0 0 1 5.94 1.82a5.8 5.8 0 0 1 2.68 2a5.64 5.64 0 0 1 .63 2.11c.64 3.13 1.31 9.28 1.5 10.76v16.88h-2.69v2.69a2.3 2.3 0 1 1-4.59 0v-2.69H15.28v2.69a2.3 2.3 0 1 1-4.59 0v-2.69H8V21.63c.16-1.48.83-7.59 1.47-10.76a5.64 5.64 0 0 1 .63-2.11a5.8 5.8 0 0 1 2.68-2a23.87 23.87 0 0 1 5.94-1.82A30.16 30.16 0 0 1 24 4.5Zm-10.54 6.85a1.41 1.41 0 0 0-1.35 1.22l-1.27 9.22A1.41 1.41 0 0 0 12 23.38h23.76A1.42 1.42 0 0 0 37.18 22a1.23 1.23 0 0 0 0-.19l-1.27-9.22a1.41 1.41 0 0 0-1.4-1.22h-21ZM13 28.81a2.41 2.41 0 1 0 2.4 2.4a2.4 2.4 0 0 0-2.4-2.4Zm22 0a2.41 2.41 0 1 0 2.41 2.4a2.39 2.39 0 0 0-2.41-2.4Zm-24.31 9.7h26.62M15.63 8.48h16.68"/>`),
		g.Group(children),
	)
}