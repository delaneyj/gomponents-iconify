package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beecount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.437 33.824l2.49 7.174l2.636-7.171Zm18.801-7.778H32.66v-7.721l7.577.013c4.378.097 4.323 7.705 0 7.707Zm0 0c3.11 2.545-.34 7.138-4.033 3.706l-3.544-3.706m7.577 0l-7.578-7.721M28.348 9.85l1.8-2.848"/><path fill="currentColor" d="M27.205 12.761a.75.75 0 1 0 .75-.75a.75.75 0 0 0-.75.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.34 18.325V32.8a1.025 1.025 0 0 0 1.025 1.025h15.27A1.025 1.025 0 0 0 32.66 32.8V18.325Zm-7.578.013l7.578-.013v7.72H7.762c-4.322-.002-4.377-7.61 0-7.707Zm7.577 7.708l-3.544 3.706c-3.694 3.432-7.143-1.161-4.033-3.707m7.578-7.72l-7.578 7.721M24 15.428h-8.66v-.001c0-3.56 3.877-6.447 8.66-6.447s8.66 2.886 8.66 6.447v0Zm-6.148-8.426l1.8 2.848"/><path fill="currentColor" d="M20.045 12.011a.75.75 0 1 0 .75.75a.75.75 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}