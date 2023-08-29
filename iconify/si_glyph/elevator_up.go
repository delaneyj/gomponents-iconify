package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.482 13.986H2.334C1.598 13.986 1 13.236 1 12.5s.598-1.44 1.334-1.44h2.143l8.041-6.043h3.148c.736 0 1.334.704 1.334 1.441c0 .735-.598 1.523-1.334 1.523h-2.184l-8 6.005z"/><g transform="translate(5 1)"><ellipse cx="1.438" cy="1.479" rx="1.438" ry="1.479"/><path d="M3.924 4.155a.327.327 0 0 0-.471-.045l-.889.771c-.201-.5-.676-.85-1.231-.85c-.736 0-1.334.618-1.334 1.384v2.584L3.88 4.644a.359.359 0 0 0 .044-.489z"/></g></g>`),
		g.Group(children),
	)
}