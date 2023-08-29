package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M30 13.5a7.47 7.47 0 0 1-3.57-.9A9.83 9.83 0 0 1 28 17.89a10 10 0 1 1-15.83-8a1.4 1.4 0 0 1 1.94.31a1.37 1.37 0 0 1-.31 1.92a7.18 7.18 0 1 0 11.43 5.8a7.07 7.07 0 0 0-3-5.76A1.37 1.37 0 0 1 22 10.2a1.38 1.38 0 0 1 1.52-.49a7.45 7.45 0 0 1-.3-6.83a16.06 16.06 0 1 0 9.93 9.93a7.46 7.46 0 0 1-3.15.69ZM16.77 8.65a1.29 1.29 0 0 1 2.58 0v9.75a1.29 1.29 0 0 1-2.58 0Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}