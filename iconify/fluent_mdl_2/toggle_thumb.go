package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleThumb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1031 0q143 0 274 36t244 103t205 161t157 208t101 245t36 273q0 140-36 270t-103 243t-159 208t-206 160t-243 104t-270 37q-140 0-271-36t-245-104t-209-160t-163-207t-105-244t-38-271q0-143 37-274t104-246t161-207t208-159T756 37t275-37z"/>`),
		g.Group(children),
	)
}