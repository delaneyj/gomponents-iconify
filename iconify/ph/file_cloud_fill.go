package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCloudFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 181a52.06 52.06 0 0 1-52 51H60.72C40.87 232 24 215.77 24 195.92a36 36 0 0 1 19.28-31.79a4 4 0 0 1 5.77 4.33a63.53 63.53 0 0 0-1 11.15a8.22 8.22 0 0 0 7.5 8.39a8 8 0 0 0 8.45-8a47.55 47.55 0 0 1 4.37-20A48 48 0 0 1 160 181Zm56-93v128a16 16 0 0 1-16 16h-24a8 8 0 0 1 0-16h24V96h-48a8 8 0 0 1-8-8V40H56v88a8 8 0 0 1-16 0V40a16 16 0 0 1 16-16h96a8 8 0 0 1 5.66 2.34l56 56A8 8 0 0 1 216 88Zm-27.31-8L160 51.31V80Z"/>`),
		g.Group(children),
	)
}