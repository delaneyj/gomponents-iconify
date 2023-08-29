package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M170.5 62Q152 62 139 75t-13 31.5t13 31.5t31.5 13t31.5-13t13-31.5T202 75t-31.5-13zm.5 192q-44 0-87 16.5T41 299v23h260v-23q0-12-43-28.5T171 254zm-.5-233Q206 21 231 46t25 60.5t-25 60.5t-60.5 25t-60.5-25t-25-60.5T110 46t60.5-25zm0 192q31.5 0 69.5 9t69.5 29.5T341 299v64H0v-64q0-27 31.5-47.5T101 222t69.5-9z"/>`),
		g.Group(children),
	)
}