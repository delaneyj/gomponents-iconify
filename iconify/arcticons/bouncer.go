package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bouncer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="15.153" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="10.669" ry="10.653"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.597 20.192a9.558 9.558 0 0 0-7.011 9.207V43.5h32.828V29.399a9.558 9.558 0 0 0-7.012-9.207m-15.186-8.154h3.899a.912.912 0 0 1 .912.912v1.549a2.002 2.002 0 0 1-2.001 2.001h-1.72a2.002 2.002 0 0 1-2.002-2.002V12.95a.912.912 0 0 1 .913-.912Zm7.669 0h3.898a.912.912 0 0 1 .913.912v1.549a2.002 2.002 0 0 1-2.002 2.001h-1.72a2.002 2.002 0 0 1-2.001-2.002V12.95a.912.912 0 0 1 .912-.912Zm-.912 1.905h-1.946"/>`),
		g.Group(children),
	)
}