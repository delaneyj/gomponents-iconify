package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HamburgerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 22V21C40 12.1634 32.8366 5 24 5C15.1634 5 8 12.1634 8 21V22"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 34C8 34 6 39 10 41C14 43 34 43 38 41C42 39 40 34 40 34"/><rect width="38" height="14" x="5" y="21" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="7"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M11 29L12.5917 28.2042C14.1235 27.4382 15.9098 27.3639 17.5 28L17.9239 28.1695C19.2444 28.6978 20.7279 28.6361 22 28V28C23.2721 27.3639 24.7556 27.3022 26.0761 27.8305L27.0353 28.2141C28.285 28.714 29.6888 28.6556 30.8927 28.0537V28.0537C32.2193 27.3904 33.7807 27.3904 35.1073 28.0537L37 29"/><circle cx="17" cy="15.05" r="2.5" fill="#000"/><circle cx="23.75" cy="12.3" r="2.5" fill="#000"/></g>`),
		g.Group(children),
	)
}