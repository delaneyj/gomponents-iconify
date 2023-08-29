package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contao(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-89q82-70 123.5-172.5t25.5-208.5q-42 10-126 31t-126 30q0 95-33 143.5t-95 48.5q-128 0-160-128q-5-21-35.5-131t-45.5-185t-15-132q0-53 45-90.5t115-37.5q49 0 98.5 40t61.5 88q38-8 76.5-19t85-25.5t62.5-19.5q-26-88-89.5-154.5T626.428 0h270q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-828-580q33 192 85 379q35 119 140 201h-165q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h187q-123 55-195.5 181t-51.5 263z"/>`),
		g.Group(children),
	)
}