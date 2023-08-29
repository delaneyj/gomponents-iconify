package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserManual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Zm4.33-2v39H37.6a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/><ellipse cx="25.531" cy="33.668" fill="currentColor" rx=".823" ry=".75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.301 20.049a6.432 6.432 0 0 1 1.959-4.763a7.274 7.274 0 0 1 5.224-1.786c3.918 0 7.183 2.977 7.183 6.548a6.432 6.432 0 0 1-1.959 4.763c-1.632 1.19-5.224 2.977-5.224 5.656"/>`),
		g.Group(children),
	)
}