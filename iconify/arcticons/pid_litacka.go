package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PidLitacka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.425 4.945C9.35 1.724 4.695 16.64 5.473 28.622C9.128 7.041 29.003 6.33 36.888 13.41c9.086 8.156 10.086 32.131-17.7 29.952"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.854 35.626c18.192 8.964 31.221-5.454 28.523-15.303m-18.918-1.24v9.232"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.155 22.464a3.39 3.39 0 0 0-3.381-3.38h0a3.39 3.39 0 0 0-3.38 3.38v2.197a3.39 3.39 0 0 0 3.38 3.381h0a3.39 3.39 0 0 0 3.38-3.38m.001 3.38V14.52M12.794 24.855a3.39 3.39 0 0 0 3.38 3.381h0a3.39 3.39 0 0 0 3.38-3.38v-2.198a3.39 3.39 0 0 0-3.38-3.38h0a3.39 3.39 0 0 0-3.38 3.38m0-3.381V32.8"/>`),
		g.Group(children),
	)
}