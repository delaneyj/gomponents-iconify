package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vala(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M19.27 512C8.626 512 0 503.373 0 492.73V19.27C0 8.626 8.627 0 19.27 0h473.46C503.374 0 512 8.627 512 19.27v473.46c0 10.643-8.627 19.27-19.27 19.27H19.27zm248.32-409.305c-39.762-3.602-83.167 4.483-131.508 41.97c-51.75 40.133-56.773 116.458 7.406 113.07c-13.76-26.054-9.481-101.96 59.079-126.676l7.033 298.354l62.904.01l121.338-325.456h-30.494l-91.913 267.436l-3.845-268.708z"/>`),
		g.Group(children),
	)
}