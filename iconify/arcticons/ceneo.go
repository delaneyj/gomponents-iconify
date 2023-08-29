package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ceneo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M21.6 27.977v-7.954l5.27 7.954v-7.954m-13.1 5.287v.032a2.635 2.635 0 0 1-2.635 2.635h0A2.635 2.635 0 0 1 8.5 25.342v-2.684a2.635 2.635 0 0 1 2.635-2.635h0a2.635 2.635 0 0 1 2.635 2.635v.032M15.844 24h2.593m1.384 3.977h-3.977v-7.954h3.977M29.03 24h2.593m1.384 3.977H29.03v-7.954h3.977"/><rect width="5.27" height="7.954" x="34.23" y="20.023" rx="2.635" ry="2.635"/></g>`),
		g.Group(children),
	)
}