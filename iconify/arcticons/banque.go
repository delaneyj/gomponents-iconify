package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Banque(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.23 33.66h11.21M20.93 43.5h19.99m-5.41-15.41l3.69-2.21l-2.96-4.98m-1.95-5.26L17.42 42.38c-5.94-3.65-7.13-13.32-2.35-20C20.55 14.74 34.88 6.87 42.33 4.5M5.67 26.47h15.41m-10.1-8.16h15.41M8.35 22.39h15.41"/><circle cx="33.87" cy="23.86" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}