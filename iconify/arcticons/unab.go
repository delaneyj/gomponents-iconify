package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="23.178" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="16.871" ry="18.678"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.444 36.78l-4.02 2.97l-.018 3.75m27.15-6.72l4.02 2.97l.018 3.75M14.409 17.953h4.737v10.845h-4.737zm14.444 0h4.737v10.845h-4.737zm-2.483 0l-2.354-2.385l-2.385 2.385V28.8h4.738Z"/>`),
		g.Group(children),
	)
}