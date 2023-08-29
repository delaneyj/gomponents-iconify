package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152Zm-66.43-92.76a12 12 0 0 0-2.35-16.82a12 12 0 0 0-16.8 2.36a11.7 11.7 0 0 0-1.74 3.22a12 12 0 0 1-22.63-8a36.45 36.45 0 0 1 5.2-9.67a36 36 0 0 1 57.5 43.34L128 164h24a12 12 0 0 1 0 24h-48a12 12 0 0 1-9.6-19.2Z"/>`),
		g.Group(children),
	)
}