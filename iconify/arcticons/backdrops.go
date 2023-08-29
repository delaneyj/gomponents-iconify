package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backdrops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="16.608" height="31.707" x="15.696" y="8.147" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.238"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.304 13.248h5.562a1.235 1.235 0 0 1 1.237 1.238v19.027a1.235 1.235 0 0 1-1.237 1.238h-5.43m-17.06 0h-5.242a1.235 1.235 0 0 1-1.237-1.238V14.486a1.235 1.235 0 0 1 1.237-1.238h5.563"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.103 16.384h3.16a1.235 1.235 0 0 1 1.237 1.238v12.755a1.235 1.235 0 0 1-1.237 1.238h-3.16m-30.207 0H5.737A1.235 1.235 0 0 1 4.5 30.377V17.622a1.235 1.235 0 0 1 1.237-1.238h3.16"/>`),
		g.Group(children),
	)
}