package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apkupdater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39 13.71l2.59-2.55a3 3 0 0 0-4.28-4.27h0l-2.55 2.6a18.44 18.44 0 0 0-21.52 0l-2.57-2.57A3.14 3.14 0 0 0 8.5 6a3 3 0 0 0-2.12.88a3 3 0 0 0 0 4.26L9 13.7a18.55 18.55 0 0 0-3.5 10.81v15.57a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V24.51a18.51 18.51 0 0 0-3.6-10.8Zm-25.86 7a3.07 3.07 0 1 1 3.07 3.07a3.08 3.08 0 0 1-3.07-3.07ZM24 40.05a7 7 0 1 1 7-7v0a7 7 0 0 1-7 7Zm7.79-16.29a3.07 3.07 0 1 1 3.06-3.07a3.07 3.07 0 0 1-3.06 3.07ZM24 28.58v8.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.09 34.06L24 37.29l2.91-3.23"/>`),
		g.Group(children),
	)
}