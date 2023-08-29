package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Direction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.06 1.941c-.586-.586-1.144-.033-3.041.879C9.944 5.259 1.1 10.216 1.1 10.216L8.699 11.3l1.085 7.599s4.958-8.843 7.396-13.916c.912-1.898 1.465-2.456.88-3.042zm-1.824 1.955l-5.519 10.247l-.561-4.655l6.08-5.592z"/>`),
		g.Group(children),
	)
}