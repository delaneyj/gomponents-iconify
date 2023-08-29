package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JavaScriptLanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M385 1534q65 0 104-33t59-84t27-108t6-108V405h145v794q0 82-17 165t-57 149t-106 109t-163 42q-30 0-60-4t-59-17v-143q26 19 57 26t64 8zm751-832q0 46 17 81t53 64q44 36 96 62t103 52q57 29 115 62t106 78t77 100t30 130q0 96-38 160t-101 102t-141 55t-161 16q-31 0-73-5t-87-14t-84-22t-66-30v-171q29 26 70 46t87 35t92 23t84 8q49 0 96-8t85-30t61-59t24-96q0-52-25-91t-68-72t-96-59t-109-56t-110-60t-96-72t-67-93t-26-123q0-88 38-150t100-103t139-59t153-19q30 0 65 1t71 6t69 14t62 23v163q-62-42-134-59t-147-17q-44 0-90 9t-85 31t-64 58t-25 89z"/>`),
		g.Group(children),
	)
}