package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRadishOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M39 15c2.183.555 4 1.5 4 4s-3 3-5 3M26 10c0-3 1.749-5 4-5c2 0 3.048.62 4 3m-3 4s.951-2.132 2.123-2.91c1.87-1.243 4.377-1.372 5.748 0a4.057 4.057 0 0 1 0 5.748C37.28 16.43 35.997 17 35.997 17"/><path fill="#555" d="M10 14c-5.084 5.085-6.417 12.913-3.41 19.103C6.59 33.103 5 41 6 42s8.862-.592 8.862-.592A15.96 15.96 0 0 0 21.802 43c4.334 0 8.97-1.853 12.199-5c3.184-3.184 4.607-7.298 4.607-11.764c0-4.507-1.607-8.236-4.953-11.711c-3.346-3.476-7.254-5.213-11.766-5.213c-4.507 0-8.704 1.504-11.888 4.688Z"/><path d="M5 28s5-1 7 3m6 11s1-4-1-6m4.803 7a15.95 15.95 0 0 1-6.94-1.592"/><path d="M6.59 33.103C3.582 26.913 4.915 19.085 10 14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRadishOne0)"/>`),
		g.Group(children),
	)
}