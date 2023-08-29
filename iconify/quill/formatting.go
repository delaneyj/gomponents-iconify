package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Formatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.75 18L10 8L6.25 18m7.5 0L16 24m-2.25-6h-7.5M4 24l2.25-6M28 16c-.333.833-1.8 2-5 2s-3.5 2-3.5 3s.7 3 3.5 3s5-2.5 5-4.5m0-3.5v3.5m0-3.5c0-1.333-.8-4-4-4c-2.4 0-3.667 1.333-4 2m8 5.5V24"/>`),
		g.Group(children),
	)
}