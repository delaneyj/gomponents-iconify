package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsBy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M314.9 194.4v101.4h-28.3v120.5h-77.1V295.9h-28.3V194.4c0-4.4 1.6-8.2 4.6-11.3c3.1-3.1 6.9-4.7 11.3-4.7H299c4.1 0 7.8 1.6 11.1 4.7c3.1 3.2 4.8 6.9 4.8 11.3zm-101.5-63.7c0-23.3 11.5-35 34.5-35s34.5 11.7 34.5 35c0 23-11.5 34.5-34.5 34.5s-34.5-11.5-34.5-34.5zM247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8zm.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3z"/>`),
		g.Group(children),
	)
}