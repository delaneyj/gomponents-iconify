package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Users(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm.63 2.113a3 3 0 1 0-4.259 0A3.997 3.997 0 0 0 1 9.5V13a2 2 0 0 0 2 2h4c.597 0 1.134-.262 1.5-.677c.366.415.903.677 1.5.677h3a2 2 0 0 0 2-2v-2c0-1.218-.622-2.29-1.565-2.917a2.5 2.5 0 1 0-3.87 0c-.241.16-.462.35-.656.564a4.005 4.005 0 0 0-1.78-2.534ZM5 7a2.5 2.5 0 0 0-2.5 2.5V13a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V9.5A2.5 2.5 0 0 0 5 7Zm7.5-.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 2.5a2 2 0 0 0-2 2v2a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-2a2 2 0 0 0-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}