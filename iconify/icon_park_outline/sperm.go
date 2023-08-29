package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sperm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M18.237 24.475c1.856 1.299 2.33 2.674 3.609 3.57c1.4.98 2.947 1.5 4.169 1.014c2.307-.916 3.976-3.908 6.011-6.815c3.96-5.655 3.954-14.385.26-16.971c-3.692-2.586-11.843.433-15.802 6.088c-1.935 2.763-4.47 6.445-4.317 8.002c.129 1.311.57 2.042 1.958 3.275s2.132.45 4.112 1.837Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M13.617 22.317c-3.54 3.898-4.008 6.86-1.402 8.885c2.605 2.026 5.877 1.027 9.815-2.995"/><path stroke-linecap="round" d="M12.239 31.227c-3.097 3.388-3.667 6.546-1.71 9.477c2.937 4.396 8.755 4.155 11.595.879c2.84-3.277 8.184-11.396 14.059-9.727c5.875 1.669 4.877 8.088.939 8.762"/></g>`),
		g.Group(children),
	)
}