package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.9 14.273l-.055 19.454L24.027 43.5l.056-19.453m0 0L24.027 43.5L7.1 33.727l.055-19.454l16.928 9.774z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.791 18.986l8.109-4.713L23.973 4.5L7.155 14.273l16.927 9.774l4.83-2.807l2.493 1.503l1.386-3.757z"/><ellipse cx="24.028" cy="14.273" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.219" ry="1.147"/><ellipse cx="30.683" cy="14.273" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.219" ry="1.147"/><ellipse cx="17.372" cy="14.273" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.219" ry="1.147"/><ellipse cx="18.919" cy="27.762" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.231" ry="2.307" transform="rotate(-36.149 18.919 27.762)"/><ellipse cx="12.263" cy="23.513" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.231" ry="2.307" transform="rotate(-36.149 12.263 23.513)"/><ellipse cx="28.479" cy="27.762" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.307" ry="1.231" transform="rotate(-53.851 28.479 27.762)"/><ellipse cx="35.134" cy="23.513" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.307" ry="1.231" transform="rotate(-53.851 35.134 23.513)"/><ellipse cx="18.919" cy="34.261" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.231" ry="2.307" transform="rotate(-36.149 18.919 34.26)"/><ellipse cx="12.263" cy="30.012" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.231" ry="2.307" transform="rotate(-36.149 12.263 30.012)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.654 33.308c2.86 1.135 5.766-.495 7.628-4.421"/>`),
		g.Group(children),
	)
}