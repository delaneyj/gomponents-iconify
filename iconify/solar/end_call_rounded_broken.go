package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndCallRoundedBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M20.917 10.503C19.56 9.04 16.807 7 12 7c-1.16 0-2.2.119-3.13.32M22 13.85c0 2.06-1.782 3.565-3.607 3.047l-1.34-.38C15.844 16.174 15 14.983 15 13.618c0 0 0-1.654-3-1.654c-2.999 0-3 1.653-3 1.654c0 1.365-.844 2.556-2.053 2.9l-1.34.38C3.782 17.414 2 15.91 2 13.85c0-1.237.277-2.477 1.083-3.347a10.24 10.24 0 0 1 2.55-1.985"/>`),
		g.Group(children),
	)
}