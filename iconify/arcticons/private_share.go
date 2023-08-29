package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivateShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c9.043-3.117 15.489-10.363 16.5-19.589a79.36 79.36 0 0 0-.071-12.027a2.541 2.541 0 0 0-2.468-2.366c-4.091-.126-8.846-.808-12.52-4.427a2.052 2.052 0 0 0-2.881 0c-3.675 3.619-8.43 4.301-12.52 4.427a2.541 2.541 0 0 0-2.468 2.366A79.36 79.36 0 0 0 7.5 23.911C8.511 33.137 14.957 40.383 24 43.5Z"/><circle cx="17.781" cy="23.434" r="2.907" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.345" cy="18.139" r="2.907" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.39" cy="28.73" r="2.907" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.378 24.735l5.406 2.712m-.035-8.008l-5.371 2.695"/>`),
		g.Group(children),
	)
}