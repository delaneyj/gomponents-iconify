package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBackspaceOutline0"><g id="evaBackspaceOutline1"><g id="evaBackspaceOutline2" fill="currentColor"><path d="M20.14 4h-9.77a3 3 0 0 0-2 .78l-.1.11l-6 7.48a1 1 0 0 0 .11 1.37l6 5.48a3 3 0 0 0 2 .78h9.77A1.84 1.84 0 0 0 22 18.18V5.82A1.84 1.84 0 0 0 20.14 4ZM20 18h-9.63a1 1 0 0 1-.67-.26l-5.33-4.85l5.38-6.67a1 1 0 0 1 .62-.22H20Z"/><path d="M11.29 14.71a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L15.41 12l1.3-1.29a1 1 0 0 0-1.42-1.42L14 10.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42Z"/></g></g></g>`),
		g.Group(children),
	)
}