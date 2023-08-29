package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnsibleAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m242.28 99.715l125.453 297.166L178.77 254.18l63.51-154.464zM464.958 465.88L272.075 19.74c-10.957-26.127-49.027-26.512-60.374 0L0 508.22h72.135l83.897-201.508l250.12 194.451c28.351 26.743 74.49-.207 58.806-35.283z"/>`),
		g.Group(children),
	)
}