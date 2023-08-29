package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Engine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 104a8 8 0 0 0-8 8v24h-16v-16a16 16 0 0 0-16-16h-12.69L160 68.69A15.86 15.86 0 0 0 148.69 64H128V48h24a8 8 0 0 0 0-16H88a8 8 0 0 0 0 16h24v16H48a16 16 0 0 0-16 16v56H16v-24a8 8 0 0 0-16 0v64a8 8 0 0 0 16 0v-24h16v20.69A15.86 15.86 0 0 0 36.69 184L72 219.31A15.86 15.86 0 0 0 83.31 224h65.38a15.86 15.86 0 0 0 11.31-4.69L195.31 184H208a16 16 0 0 0 16-16v-16h16v24a8 8 0 0 0 16 0v-64a8 8 0 0 0-8-8Zm-40 64h-12.69a15.86 15.86 0 0 0-11.31 4.69L148.69 208H83.31L48 172.69V80h100.69L184 115.31a15.86 15.86 0 0 0 11.31 4.69H208Z"/>`),
		g.Group(children),
	)
}