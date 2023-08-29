package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trebleshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.12 19.53v1.18a14.93 14.93 0 0 0 11.01 14.42M32.47 8.42a14.93 14.93 0 0 0-19 1.74M30.41 34.2a14.93 14.93 0 0 0 8-17.35"/><circle cx="24" cy="35.64" r="3.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.07" cy="13.25" r="3.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.93" cy="13.25" r="3.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.54 39.31a6.58 6.58 0 0 0 7.84 2.46a6.32 6.32 0 0 0 1.28-.68a6.57 6.57 0 0 0-7.33-10.91m22.42-13.89A6.57 6.57 0 0 0 31.91 9a6.72 6.72 0 0 0-.81 1.21a6.57 6.57 0 0 0 2.79 8.87M10.65 6.69a6.57 6.57 0 0 0-6.07 5.53a7 7 0 0 0-.07 1.45a6.57 6.57 0 1 0 13.12-.84"/>`),
		g.Group(children),
	)
}