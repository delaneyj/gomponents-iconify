package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.61 4.5h8.78a4.39 4.39 0 0 1 4.39 4.39v2.2h2.56A1.61 1.61 0 0 1 37 12.71v29.17a1.61 1.61 0 0 1-1.62 1.62H12.66A1.61 1.61 0 0 1 11 41.88V12.71a1.61 1.61 0 0 1 1.62-1.62h2.56v-2.2a4.39 4.39 0 0 1 4.43-4.39Zm-4.06 27.06l7.86 4.53l9-15.66m-17.19-9.34h17.56"/>`),
		g.Group(children),
	)
}