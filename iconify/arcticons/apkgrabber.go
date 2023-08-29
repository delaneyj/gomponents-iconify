package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apkgrabber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="m34.65 15.16l2.73-3.52a1 1 0 0 0-.05-1.31C26.39-1.22 5.74 5.76 5.12 23.3a.4.4 0 0 1-.08.23a.44.44 0 0 1-.37.21H2.74a.25.25 0 0 0-.19.4l5.55 7.18a.26.26 0 0 0 .39 0l5.56-7.17a.25.25 0 0 0-.2-.4h-1.93a.41.41 0 0 1-.32-.15v-.06a.42.42 0 0 1-.1-.28a12.51 12.51 0 0 1 21.67-8a1 1 0 0 0 1.48-.1Zm-21.3 17.68l-2.73 3.52a1 1 0 0 0 .05 1.31c10.94 11.55 31.59 4.57 32.21-13a.4.4 0 0 1 .08-.23a.44.44 0 0 1 .37-.21h1.93a.25.25 0 0 0 .19-.4l-5.55-7.15a.26.26 0 0 0-.39 0L34 23.85a.25.25 0 0 0 .2.4h1.93a.41.41 0 0 1 .32.15v.06a.42.42 0 0 1 .1.28a12.51 12.51 0 0 1-21.67 8a1 1 0 0 0-1.53.1Z"/>`),
		g.Group(children),
	)
}