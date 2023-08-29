package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleMaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.897 33.398a39.396 39.396 0 0 1 3.073 4.532a23 23 0 0 1 1.683 4.285c.352.99.67 1.285 1.352 1.285c.744 0 1.081-.502 1.342-1.28a23.224 23.224 0 0 1 1.637-4.2a47.876 47.876 0 0 1 4.555-6.458a41.353 41.353 0 0 0 4.511-6.5a15.78 15.78 0 0 0 1.555-6.888a13.593 13.593 0 0 0-1.533-6.358"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.735 24.578c1.457 3.329 4.267 6.255 6.168 8.822l10.099-11.962a5.306 5.306 0 0 1-4.004 1.861a5.173 5.173 0 0 1-5.197-5.19a5.505 5.505 0 0 1 1.195-3.348m8.118-9.646a13.519 13.519 0 0 1 7.961 6.7l-8.07 9.617a5.606 5.606 0 0 0 1.194-3.361a5.223 5.223 0 0 0-5.189-5.175a5.418 5.418 0 0 0-4.011 1.858m-6.414-5.389A13.478 13.478 0 0 1 23.972 4.5a13.856 13.856 0 0 1 4.134.619l-8.114 9.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.735 24.578a15.376 15.376 0 0 1-1.34-6.428a13.516 13.516 0 0 1 3.19-8.785l6.411 5.395Z"/>`),
		g.Group(children),
	)
}