package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitbucket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1.036 1.615A1.026 1.026 0 0 0 0 2.625c0 .063.005.12.016.177l4.349 26.417A1.39 1.39 0 0 0 5.73 30.38H26.6c.51.005.943-.359 1.026-.859l4.359-26.708a1.017 1.017 0 0 0-.844-1.172a1.119 1.119 0 0 0-.177-.016zm18.323 19.088h-6.661l-1.802-9.417h10.078z"/>`),
		g.Group(children),
	)
}