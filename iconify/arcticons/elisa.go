package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elisa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><rect width="32" height="19" x="8" y="12.245" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="24.617" height="6.818" x="11.692" y="20.591" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.409"/><circle cx="15.101" cy="24" r="3.409" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.212 20.591a5.32 5.32 0 0 1 0 6.819"/><circle cx="32.9" cy="24" r="3.409" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.5 39.5l-2.678-5.63H9.178L6.5 39.501"/>`),
		g.Group(children),
	)
}