package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndianoilOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 21.632h19.635m-1.079 5.212v-5.212m-4.897 5.212v-5.212m-8.807 5.212v-5.212c0-1.213.854-1.617 2.202-1.617s2.022.36 2.022 3.145c-1.438 0-2.441-.102-2.441.827c0 1.21 2.785.446 2.785 1.66c0 1.093-1.617 1.752-2.8.718m3.849-4.733a1.385 1.385 0 0 1 1.213 1.453c0 .824-.884.884-.884 1.723c0 .583.465.838 1.378.838a1.702 1.702 0 0 0 1.483-1.034m14.812-2.98a1.385 1.385 0 0 1 1.213 1.453c0 .824-.883.884-.883 1.723c0 .583.464.838 1.378.838a1.702 1.702 0 0 0 1.482-1.034m-16.16.001c.69-.69 2.217-.674 3.055-.674M7.17 21.632a9.471 9.471 0 0 1 .1 1.528c-1.437 0-2.44-.102-2.44.827c0 1.21 2.785.446 2.785 1.66c0 1.093-1.617 1.752-2.8.719a19.657 19.657 0 0 1 1.572 1.692m21.896-6.426H43.5m-1.138 5.25v-5.25m-6.201 5.25v-5.25m-4.822 5.25v-5.25m-1.947 5.25v-5.25m12.97 1.816a18.794 18.794 0 0 0-2.516 1.554"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.917 26.65c-.929-.69-1.273-2.172-.599-2.846s1.77-.524 2.592.472m-15.798-1.79a1.177 1.177 0 0 1 1.168-.854c.648 0 .989.674.989 1.37s-.787.921-1.303 1.056a1.309 1.309 0 0 1 1.304 1.236c0 .719-.338 1.355-1.192 1.355s-1.438-1.003-1.438-1.003m4.752-1.782h-.884m.704-3.849a2.15 2.15 0 0 0 2.351.075"/><circle cx="7.142" cy="20.097" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.72 17.075h40.56M3.72 30.925h40.56"/>`),
		g.Group(children),
	)
}