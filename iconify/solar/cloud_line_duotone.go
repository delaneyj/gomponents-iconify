package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14.381 9.027a5.766 5.766 0 0 1 1.905-.32a5.76 5.76 0 0 1 1.87.308C20.392 9.78 22 11.881 22 14.353C22 17.472 19.442 20 16.286 20h-10C3.919 20 2 18.104 2 15.765c0-2.34 1.919-4.235 4.286-4.235a4.291 4.291 0 0 1 2.381.713"/><path d="M7.116 11.609a5.577 5.577 0 0 1-.354-1.962C6.762 6.528 9.32 4 12.476 4c2.94 0 5.361 2.194 5.68 5.015" opacity=".5"/></g>`),
		g.Group(children),
	)
}