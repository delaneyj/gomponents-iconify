package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stocks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m990.193 384l-70-69l-284 285q-5 14-14 23q-20 20-49 16q-26 1-44-17q-9-8-13-21l-37-37q-2 3-54.5 60.5t-57.5 62.5q-20 20-50 16q-26 1-44-17q-9-8-13-21l-5-5l-122 122q-1 1-2.5 2t-2.5 2v110h832q26 0 45 19t19 45.5t-19 45t-45 18.5h-896q-27 0-45.5-18.5T.193 960V64q0-26 18.5-45t45-19t45.5 19t19 45v555l68-68q4-12 13-21q18-18 44-17q29-3 49 16q9 9 14 24l3 3l101-101q4-12 13-21q18-18 44-17q30-3 49 16q9 9 14 24l35 35l261-261l-68-68q-1-36 22-35h202q14 1 22.5 9.5t9.5 22.5v202q1 24-34 22z"/>`),
		g.Group(children),
	)
}