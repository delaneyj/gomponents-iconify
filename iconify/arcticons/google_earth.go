package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleEarth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.458 32.963c-.748-3.332 1.208-7.124 4.909-7.124c8.965 0 10.983 14.662 22.9 14.662c5.887 0 9.773-4.795 9.773-4.795"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.504 24.003a9.096 9.096 0 0 1 9.17-9.278c12.237 0 12.924 18.549 25.625 18.549c4.448 0 6.742-2.695 7.914-5.755"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.384 10.355a12.227 12.227 0 0 1 8.637-3.863c12.716 0 15.587 18.674 25.035 18.674a6.056 6.056 0 0 0 4.425-2.083"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.7 18.168a6.397 6.397 0 0 1-1.905.281c-7.038 0-7.15-15.303-21.364-15.303a20.264 20.264 0 0 0-4.68.607"/>`),
		g.Group(children),
	)
}