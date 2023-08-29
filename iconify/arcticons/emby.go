package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.25 3.5L11.12 14.63l1.75 1.75l-9.37 9.37l11.13 11.13l1.75-1.75l9.37 9.37l11.13-11.13l-1.75-1.75l9.37-9.37l-11.13-11.13l-1.75 1.75Zm-3 12.81l6.86 3.84L33 24l-6.85 3.85l-6.86 3.84V16.31Z"/>`),
		g.Group(children),
	)
}