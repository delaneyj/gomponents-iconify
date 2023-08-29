package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waveup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.372 22.21l-8.15-7.6c-2.62-1.892-5.95 1.742-3.573 3.648c0 0 6.285 6.23 8.57 8.563c1.744 1.782 5.136 5.436 5.136 5.436c-2.874-.86-5.282-2.091-6.944-1.986c-2.439.155-5.902 3.07-1.919 4.53l12.226 4.483c3.905 1.727 7.489.05 10.405-2.363c6.433-5.321 6.691-12.709 2.906-16.63c-4.66-4.83-9.22-9.53-14.223-13.946c-2.37-2.093-5.573 1.664-3.277 3.713l5.634 5.255l-8.933-8.346c-2.347-2.1-5.792 1.364-3.481 3.484l7.86 7.33l-9.295-8.633c-2.17-1.973-5.096 1.557-3.298 3.406Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.575 5.85a3.949 3.949 0 0 0-1.188-.198H9.45A3.953 3.953 0 0 0 5.5 9.604v28.792a3.953 3.953 0 0 0 3.951 3.952h12.884a3.954 3.954 0 0 0 3.889-3.245M5.5 36.955h14.866M5.5 11.066h4.96"/>`),
		g.Group(children),
	)
}