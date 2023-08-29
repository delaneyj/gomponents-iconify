package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M32 44C32 44 32 40.1355 32.0015 40.1282C32.6987 39.7817 33.3702 39.3913 34.0124 38.9606C34.9896 38.3053 35.899 37.5569 36.7279 36.7279C39.9853 33.4706 42 28.9706 42 24H6C6 28.9706 8.01472 33.4706 11.2721 36.7279C12.1086 37.5645 13.0271 38.319 14.0145 38.9786C14.653 39.4051 15.3204 39.792 16.0131 40.1355L16 44H32Z"/><path d="M24 18.0083V8"/><path d="M36 18.0083V12"/><path d="M12 18.0083V12"/><path d="M40 8C37.7909 8 36 9.79086 36 12"/><path d="M28 4C25.7909 4 24 5.79086 24 8"/><path d="M16 8C13.7909 8 12 9.79086 12 12"/></g>`),
		g.Group(children),
	)
}