package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Willyweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="15.526" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="19.835" ry="6.786" transform="rotate(-11.238 24 15.526)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.02 22.777a3.666 3.666 0 0 0-.358 2.09c.515 3.713 7.247 5.847 15.037 4.767s13.688-4.964 13.173-8.677a3.899 3.899 0 0 0-1.331-2.34"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.836 29.416a5.12 5.12 0 0 0 .047.589c.515 3.712 4.983 6.16 9.98 5.468a11.83 11.83 0 0 0 3.63-1.105"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.624 33.258a4.852 4.852 0 0 0-.07 1.712c.398 2.869 3.315 4.932 6.95 5.203"/>`),
		g.Group(children),
	)
}