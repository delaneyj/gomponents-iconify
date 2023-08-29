package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.094 22.36l.083 19.898l1.256 1.242h3.32l1.535-1.242V28.72l9.7 14.298h4.952l.967-1.327V22.36H28.69v10.37l-6.636-10.37h-8.961Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M19.556 15.55a1.53 1.53 0 1 1 0-3.06a1.53 1.53 0 0 1 0 3.06h0Zm8.905 0a1.53 1.53 0 1 1 1.53-1.53h0a1.53 1.53 0 0 1-1.53 1.53Z"/><path d="M24 6.13h0c6.024 0 10.907 4.883 10.907 10.907v1.556H13.094v-1.556C13.094 11.013 17.977 6.13 24 6.13ZM14.582 4.5l3.098 3.647M33.418 4.5L30.32 8.147"/></g>`),
		g.Group(children),
	)
}