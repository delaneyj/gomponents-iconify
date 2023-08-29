package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="221.912" cy="66.088" r="34.088" fill="currentColor"/><path fill="currentColor" d="m460.12 360.478l-47.943 11.985L393 282.971A24.126 24.126 0 0 0 369.533 264h-88.705l-6.462-56H384v-32H270.674l-4.134-35.826a24 24 0 0 0-26.593-21.091l-39.736 4.585L220.1 296h142.97l24.758 115.537l80.057-20.015Z"/><path fill="currentColor" d="M224 448a120 120 0 0 1-45.248-231.135l-3.779-32.75C115.143 204.558 72 261.334 72 328c0 83.813 68.187 152 152 152a152.06 152.06 0 0 0 130.044-73.378L344 360c-16 48-61.4 88-120 88Z"/>`),
		g.Group(children),
	)
}