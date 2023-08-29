package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsChanger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.975 24.962V14.117h1.762a5.439 5.439 0 0 1 5.423 5.423h0a5.439 5.439 0 0 1-5.423 5.423Zm11.432 0V14.117l7.186 10.845V14.117m4.383 9.625a3.232 3.232 0 0 0 2.711 1.22h1.627a2.72 2.72 0 0 0 2.711-2.711h0a2.72 2.72 0 0 0-2.711-2.711H34.55a2.72 2.72 0 0 1-2.711-2.712h0a2.72 2.72 0 0 1 2.711-2.711h1.627a2.911 2.911 0 0 1 2.712 1.22M15.567 33.883h1.611m-1.611-5.369h1.611m-.806 0v5.369m2.971 0v-5.369h1.745a1.812 1.812 0 1 1 0 3.624h-1.745m7.871-1.812l-1.343 3.557l-1.342-3.557"/><circle cx="30.621" cy="32.071" r="1.812" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.232 29.118a1.648 1.648 0 0 0-1.477-.67h-.134a1.804 1.804 0 0 0-1.812 1.811v1.812"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}