package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Zm-16.54-.917a6.59 6.59 0 0 1 6.55-5.833a6.585 6.585 0 0 1 5.242 2.592a.75.75 0 0 1-1.192.911a5.085 5.085 0 0 0-4.05-2.003a5.09 5.09 0 0 0-5.037 4.333h.364a.75.75 0 0 1 .53 1.281l-1.169 1.167a.75.75 0 0 1-1.06 0L4.47 12.364a.75.75 0 0 1 .53-1.28h.46Zm12.902-.614a.75.75 0 0 0-1.06 0l-1.168 1.167a.75.75 0 0 0 .53 1.28h.363a5.09 5.09 0 0 1-5.037 4.334a5.084 5.084 0 0 1-4.037-1.986a.75.75 0 0 0-1.188.916a6.584 6.584 0 0 0 5.226 2.57a6.59 6.59 0 0 0 6.549-5.833H19a.75.75 0 0 0 .53-1.281l-1.168-1.167Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}