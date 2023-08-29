package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoSucKhoeDienTuAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.41 37.523c.49.639 1.106.876 1.962.876h1.184a1.996 1.996 0 0 0 1.996-1.995v-.009a1.996 1.996 0 0 0-1.996-1.996H9.249a1.998 1.998 0 0 1-1.997-1.997h0c0-1.106.896-2.002 2.002-2.002h1.178c.856 0 1.47.237 1.961.876m1.767 6.247c.49.639 1.106.876 1.962.876h1.184a1.996 1.996 0 0 0 1.996-1.995v-.009a1.996 1.996 0 0 0-1.996-1.996h-1.307a1.998 1.998 0 0 1-1.997-1.997h0c0-1.106.896-2.002 2.002-2.002h1.178c.856 0 1.47.237 1.961.876m2.057-.877v8m4.3 0l-3.29-4l3.29-3.973m-3.29 3.973H21.2m7.05 4.027v-8h1.8a3.5 3.5 0 0 1 3.5 3.5v1a3.5 3.5 0 0 1-3.5 3.5h-1.8Zm-1.3-4h2.6m5.45-4h5.3m-2.65 8v-8M30.449 8.592a3.373 3.373 0 0 0-2.62 5.503l5.41 6.35l5.357-6.287l.026-.03l.027-.033a3.377 3.377 0 1 0-5.41-4.027a3.367 3.367 0 0 0-2.787-1.477h-.003Z"/>`),
		g.Group(children),
	)
}