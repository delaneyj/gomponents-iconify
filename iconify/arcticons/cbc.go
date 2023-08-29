package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cbc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="7.986" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.126 17.123a6.877 6.877 0 0 0 0 13.754V17.123Zm17.748 13.754a6.877 6.877 0 0 0 0-13.754v13.754Zm6.876-15.751a6.877 6.877 0 0 0-6.876-6.877v6.877h6.877ZM32.874 39.75a6.877 6.877 0 0 0 6.877-6.876h-6.877v6.877ZM8.25 15.126a6.877 6.877 0 0 1 6.876-6.877v6.877H8.25Zm6.876 24.624a6.877 6.877 0 0 1-6.877-6.876h6.877v6.877Zm15.751-24.624a6.877 6.877 0 0 0-13.754 0h13.754ZM17.123 32.874a6.877 6.877 0 0 0 13.754 0H17.123Zm13.754 8.652H17.123A7.94 7.94 0 0 0 24 45.5a7.94 7.94 0 0 0 6.877-3.974ZM17.123 6.474h13.754A7.94 7.94 0 0 0 24 2.5a7.94 7.94 0 0 0-6.877 3.974Zm24.403 10.649v13.754A7.94 7.94 0 0 0 45.5 24a7.94 7.94 0 0 0-3.974-6.877ZM6.474 30.877V17.123A7.94 7.94 0 0 0 2.5 24a7.94 7.94 0 0 0 3.974 6.877Z"/>`),
		g.Group(children),
	)
}