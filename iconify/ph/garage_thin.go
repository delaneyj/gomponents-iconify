package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 196h-12V98.67a12 12 0 0 0-5.34-10l-88-58.67a12 12 0 0 0-13.32 0l-88 58.67a12 12 0 0 0-5.34 10V196H16a4 4 0 0 0 0 8h224a4 4 0 0 0 0-8ZM36 98.67a4 4 0 0 1 1.78-3.33l88-58.66a4 4 0 0 1 4.44 0l88 58.66a4 4 0 0 1 1.78 3.33V196h-32v-60a4 4 0 0 0-4-4H72a4 4 0 0 0-4 4v60H36ZM180 140v24h-48v-24Zm-56 24H76v-24h48Zm-48 8h48v24H76Zm56 0h48v24h-48Z"/>`),
		g.Group(children),
	)
}