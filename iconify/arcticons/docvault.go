package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Docvault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c9.043-3.117 15.489-10.363 16.5-19.589c.28-4.005.257-8.025-.071-12.027a2.541 2.541 0 0 0-2.468-2.366c-4.091-.126-8.846-.808-12.52-4.427a2.052 2.052 0 0 0-2.881 0c-3.675 3.619-8.43 4.301-12.52 4.427a2.541 2.541 0 0 0-2.468 2.366A79.362 79.362 0 0 0 7.5 23.911C8.512 33.137 14.957 40.383 24 43.5Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="21.196" r="3.458"/><path d="M24 28.262v-3.608"/></g>`),
		g.Group(children),
	)
}