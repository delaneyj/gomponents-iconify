package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gym(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.274 9.869l-3.442-4.915l1.639-1.147l3.441 4.915l-1.638 1.147Zm-1.884 2.54L16.67 9.95l-8.192 5.736l1.72 2.457l-1.638 1.148l-4.588-6.554L5.61 11.59l1.72 2.458l8.192-5.736l-1.72-2.458l1.638-1.147l4.588 6.553l-1.638 1.148Zm2.375-5.326l1.638-1.147l-1.147-1.638l-1.638 1.147l1.147 1.638ZM7.168 19.046l-3.442-4.915l-1.638 1.147l3.441 4.915l1.639-1.147Zm-2.786-.491l-1.638 1.147l-1.147-1.638l1.638-1.147l1.147 1.638Z"/>`),
		g.Group(children),
	)
}