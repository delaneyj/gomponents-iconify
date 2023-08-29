package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M8 19.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm7.5-5.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11Z"/><path d="M15.509 31h-.2a11.5 11.5 0 0 1-4.6-21.956a.355.355 0 0 0 .209-.32v-.083l-1.33-2.95a1.7 1.7 0 0 1 .253-1.822l2.114-2.264A1.7 1.7 0 0 1 13.262 1h4.65a1.768 1.768 0 0 1 1.3.578l1.926 2.3a1.76 1.76 0 0 1 .248 1.806l-1.3 2.891v.151a.343.343 0 0 0 .2.31A11.557 11.557 0 0 1 27 19.5A11.517 11.517 0 0 1 15.509 31ZM12.823 8l.095.211v.515a2.363 2.363 0 0 1-1.375 2.138a9.492 9.492 0 1 0 7.905-.011a2.344 2.344 0 0 1-1.365-2.127v-.581L18.148 8h-5.325ZM11.52 5h7.949L17.8 3h-4.411l-1.868 2Z"/></g>`),
		g.Group(children),
	)
}