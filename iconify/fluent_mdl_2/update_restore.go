package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpdateRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0q141 0 272 36t245 103t207 160t160 208t103 245t37 272q0 141-36 272t-103 245t-160 207t-208 160t-245 103t-272 37q-172 0-330-55t-289-154t-226-238t-141-304l123-34q40 145 123 265t198 208t253 135t289 49q123 0 237-32t214-90t182-141t140-181t91-214t32-238q0-123-32-237t-90-214t-141-182t-181-140t-214-91t-238-32q-129 0-251 36T546 267T355 428T215 640h297v128H0V256h128v274q67-123 163-221t212-166T752 37t272-37z"/>`),
		g.Group(children),
	)
}