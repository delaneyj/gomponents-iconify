package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calibrecompanion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Zm4.331-2v39h24.87a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.396 27.201a1.224 1.224 0 1 1 1.224-1.223a1.223 1.223 0 0 1-1.224 1.223Zm7.121 0a1.224 1.224 0 1 1 1.223-1.223a1.224 1.224 0 0 1-1.223 1.223Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.95 19.669h0a8.721 8.721 0 0 1 8.72 8.721v0a1.245 1.245 0 0 1-1.244 1.245H18.473a1.245 1.245 0 0 1-1.245-1.245v0a8.721 8.721 0 0 1 8.722-8.721Zm-7.532-1.304l2.478 2.917m12.585-2.917l-2.478 2.917"/>`),
		g.Group(children),
	)
}