package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c9.043-3.117 15.489-10.363 16.5-19.589c.28-4.005.257-8.025-.071-12.027a2.541 2.541 0 0 0-2.468-2.366c-4.091-.126-8.846-.808-12.52-4.427a2.052 2.052 0 0 0-2.881 0c-3.675 3.619-8.43 4.301-12.52 4.427a2.541 2.541 0 0 0-2.468 2.366A79.362 79.362 0 0 0 7.5 23.911C8.512 33.137 14.957 40.383 24 43.5Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.322 21.048h11.356c.43 0 .78.349.78.78v8.392c0 .43-.35.78-.78.78H18.322a.78.78 0 0 1-.78-.78v-8.393c0-.43.35-.78.78-.78h0Z"/><path d="M19.726 21.048v-1.79a4.27 4.27 0 0 1 8.541 0v1.79"/><circle cx="23.996" cy="26.024" r="1.656"/></g>`),
		g.Group(children),
	)
}