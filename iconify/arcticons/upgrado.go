package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upgrado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.313 21.744l29.325.023l-1.989 4.091l-29.385.093l2.05-4.207ZM6.547 30.82l29.327.022l-1.99 4.091l-29.384.093l2.047-4.206Zm7.628-17.846l29.325.023l-1.988 4.09l-29.387.093l2.05-4.206Z"/>`),
		g.Group(children),
	)
}