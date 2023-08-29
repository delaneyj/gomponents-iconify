package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videocamerathree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M985 571L768 386v126q0 53-37.5 90.5T640 640H128q-53 0-90.5-37.5T0 512V128q0-53 37.5-90.5T128 0h512q53 0 90.5 37.5T768 128v126L985 69q16-13 39 7v488q-23 20-39 7zM576.5 128q-26.5 0-45.5 18.5t-19 45t19 45.5t45.5 19t45-19t18.5-45.5t-18.5-45t-45-18.5z"/>`),
		g.Group(children),
	)
}