package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 154a29.87 29.87 0 0 0-19.5 7.23l-25.62-19.93A29.83 29.83 0 0 0 158 128a30.52 30.52 0 0 0-.22-3.6L174 119a30 30 0 1 0-4-15a30.52 30.52 0 0 0 .22 3.6L154 113a29.91 29.91 0 0 0-32.42-14.31l-8.14-18.3a30 30 0 1 0-11 4.88l8.14 18.3a29.92 29.92 0 0 0-8.52 39.43L74 168a30.08 30.08 0 1 0 8 9l28-25a29.91 29.91 0 0 0 37.47-1.23l25.62 19.93A30 30 0 1 0 200 154Zm0-68a18 18 0 1 1-18 18a18 18 0 0 1 18-18ZM78 56a18 18 0 1 1 18 18a18 18 0 0 1-18-18ZM56 210a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm72-64a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm72 56a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}