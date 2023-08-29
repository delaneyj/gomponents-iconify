package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Saveto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 10.18L10 5.5h30.5a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.92 5.5h22.16v10.25H12.92zm0 21.25h22.16V42.5H12.92zm18.72-14.51V9.01M17.79 31.93h12.42m-12.42 5.39h12.42"/>`),
		g.Group(children),
	)
}