package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emberdotjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v24h24V0H0zm12.29 4.38c1.66-.03 2.83.42 3.84 1.85c2.25 5.58-6 8.4-6 8.4s-.23 1.48 2.02 1.42c2.78 0 5.7-2.15 6.81-3.06a.66.66 0 0 1 .9.05l.84.87a.66.66 0 0 1 .01.9c-.72.8-2.42 2.46-4.97 3.53c0 0-4.26 1.97-7.13.1a4.95 4.95 0 0 1-2.38-3.83s-2.08-.11-3.42-.63c-1.33-.52.01-2.1.01-2.1s.42-.65 1.2 0s2.24.36 2.24.36c.13-1.03.35-2.38.98-3.81c1.34-3 3.38-4.01 5.05-4.05zm.33 2.8c-1.1.07-2.8 1.78-2.88 4.93c0 0 .75.23 2.41-.91c1.67-1.14 2-2.97 1.11-3.81a.82.82 0 0 0-.64-.21Z"/>`),
		g.Group(children),
	)
}