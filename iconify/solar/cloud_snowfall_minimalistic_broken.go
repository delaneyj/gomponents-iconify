package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnowfallMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="18" r="1" fill="currentColor"/><circle cx="12" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M16 19.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-6 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 13.353c0 2.343-1.444 4.353-3.5 5.207M14.381 8.027a5.765 5.765 0 0 1 1.905-.321c.654 0 1.283.109 1.87.309m-11.04 2.594a4.351 4.351 0 0 0-.83-.08C3.919 10.53 2 12.426 2 14.765c0 1.896 1.261 3.501 3 4.041m2.116-8.197a5.577 5.577 0 0 1-.354-1.962C6.762 5.528 9.32 3 12.476 3c2.94 0 5.361 2.194 5.68 5.015m-11.04 2.594a4.29 4.29 0 0 1 1.55.634m9.49-3.228A5.73 5.73 0 0 1 20.463 9.5"/></g>`),
		g.Group(children),
	)
}