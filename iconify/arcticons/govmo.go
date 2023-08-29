package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Govmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="32.37" r="2.27" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.39 12.48c-3.63.06-11 4.15-12.78 14.55M42.5 15.15c-.28 6.31-2.73 14.67-11.08 18.19"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.61 27a53.43 53.43 0 0 0 1.48-10.91c0-5-3.3-10.8-3.3-10.8M7.61 12.48c3.63.06 11 4.15 12.78 14.55M5.5 15.15c.28 6.31 2.73 14.67 11.08 18.19"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.39 27a53.43 53.43 0 0 1-1.48-10.91c0-5 3.3-10.8 3.3-10.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.58 33.34c-3.52-1.71-4.32-6-3.75-10.63c0 0 5.06-1.42 7.56 4.32A14.71 14.71 0 0 1 24 17.54A14.71 14.71 0 0 1 27.61 27c2.5-5.74 7.56-4.32 7.56-4.32c.57 4.66-.23 8.92-3.75 10.63M7.61 42.68h32.78M6.62 38.21h34.76"/>`),
		g.Group(children),
	)
}