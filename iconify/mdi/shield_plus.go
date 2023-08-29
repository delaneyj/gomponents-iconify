package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 20v2.97h-2V20h-3v-2h3v-3h2v3h3v2h-3M12 1l9 4v6c0 .9-.1 1.78-.29 2.65A5.8 5.8 0 0 0 18 13a6 6 0 0 0-6 6c0 1.36.45 2.62 1.22 3.62L12 23c-5.16-1.26-9-6.45-9-12V5l9-4Z"/>`),
		g.Group(children),
	)
}