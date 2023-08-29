package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Animexstream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.303 32.1l6.347-7.737l-8.16-10.638l6.27-.081l5.062 6.528L35.231 5.665m-22.816 36.75L23.974 27.91l5.364 6.608l6.195-.08l-8.386-10.477l6.573-7.979m1.51-10.317l-18.13-.16m-4.685 36.91l19.19.08"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.099 5.504a20.335 20.335 0 0 0-6.713 2.994C5.09 12.456 4.74 18.47 6.336 23.35c1.647 5.032 5.583 8.884 7.967 8.75m17.301 10.396c7.03-2.09 9.602-6.172 10.516-10.13a14.3 14.3 0 0 0-8.4-16.385"/>`),
		g.Group(children),
	)
}