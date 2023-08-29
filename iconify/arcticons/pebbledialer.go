package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pebbledialer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41 5H7a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h34a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.87 28.1a40.4 40.4 0 0 1-4.1-.72a2 2 0 0 0-2 .39c-.34.34-1.33 1.33-2.37 2.33a18.59 18.59 0 0 1-8.51-8.5c1-1 2-2 2.34-2.37a2 2 0 0 0 .37-2c-.3-1.35-.55-2.72-.72-4.1A1.33 1.33 0 0 0 18.42 12H12.9a1 1 0 0 0-.88.91c-.39 5 2.42 10.35 3 11.36h0a1.09 1.09 0 0 0 .08.15h0a23.26 23.26 0 0 0 8.45 8.44h0l.3.16h0c1.27.69 6.42 3.31 11.25 2.93a1 1 0 0 0 .91-.88v-5.46a1.33 1.33 0 0 0-1.14-1.51Z"/>`),
		g.Group(children),
	)
}