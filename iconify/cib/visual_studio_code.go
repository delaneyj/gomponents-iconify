package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualStudioCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.865 3.448L24.282.281a1.99 1.99 0 0 0-2.276.385L9.397 12.171L3.902 8.004a1.33 1.33 0 0 0-1.703.073L.439 9.681a1.33 1.33 0 0 0-.005 1.969L5.2 15.999L.434 20.348a1.33 1.33 0 0 0 .005 1.969l1.76 1.604a1.33 1.33 0 0 0 1.703.073l5.495-4.172l12.615 11.51a1.982 1.982 0 0 0 2.271.385l6.589-3.172a1.993 1.993 0 0 0 1.13-1.802V5.248c0-.766-.443-1.469-1.135-1.802zm-6.86 19.818L14.432 16l9.573-7.266z"/>`),
		g.Group(children),
	)
}