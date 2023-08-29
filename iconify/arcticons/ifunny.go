package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ifunny(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38 26.87a14.24 14.24 0 0 1-27.89 0"/><ellipse cx="17.2" cy="18.44" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.5" ry="4.22"/><ellipse cx="30.61" cy="18.39" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.5" ry="4.22"/>`),
		g.Group(children),
	)
}