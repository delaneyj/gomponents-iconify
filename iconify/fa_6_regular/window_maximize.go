package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowMaximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M.3 89.5C.1 91.6 0 93.8 0 96v320c0 35.3 28.7 64 64 64h384c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64H64c-2.2 0-4.4.1-6.5.3c-9.2.9-17.8 3.8-25.5 8.2c-10.2 6-18.6 14.6-24.3 25c-3.9 7.3-6.5 15.4-7.4 24zM48 224h416v192c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16V224z"/>`),
		g.Group(children),
	)
}