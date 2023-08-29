package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carrot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1017.5 1017.5q-21.5 21.5-113-26T698.5 862T467 689T270.5 532T176 452q-22-24-14-66.5t42-90.5q-85-39-204-39v-64q94 0 170 25L0 47L47 0l170 171Q192 94 192 0h64q0 119 39 204q48-34 90.5-42t66.5 14q14 15 80 94.5T689 467t173 231.5t129.5 206t26 113z"/>`),
		g.Group(children),
	)
}