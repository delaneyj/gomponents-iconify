package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Busbahnbim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-2 11.75h13.375"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.875 17.25q.545 2.24 1 4.5c.2.994.388 1.992.5 3a20.357 20.357 0 0 1 0 4.5m0 0h-8.439"/><circle cx="10" cy="28.75" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.063 29.25H5.5m13.648-10H5.5m11.5 0l-1.125 7m4.5 0h-4.5m.5-3.5H10m9.115-4.5H23m-2.625 11H29m1-1a1 1 0 0 1-1 1m1-1q-.006-.5 0-1c.004-.333.012-.667 0-1a8.42 8.42 0 0 0-7-8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.689 20.75H25l-.75 4.5h5.653m-3.359-5.5h6.097M29 29.25h9.5m-6-9.5a6.952 6.952 0 0 1 2 .5a8.74 8.74 0 0 1 1 .5a10.028 10.028 0 0 1 2.5 2a17.232 17.232 0 0 1 2 3a6.98 6.98 0 0 1 .5 1a2.333 2.333 0 0 1 .091 1.517a1.539 1.539 0 0 1-.405.65a1.26 1.26 0 0 1-.686.333a2.87 2.87 0 0 1-.5 0q-.25-.011-.5 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.518 22.25H32.5l-.5 3h7.714"/>`),
		g.Group(children),
	)
}