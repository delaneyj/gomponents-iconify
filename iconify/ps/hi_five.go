package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HiFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m446 164l4-81H295l-23 126h1q-2 5-4 13v20h15v29q28 30 80 30q43 0 70.5-25t27.5-65q0-27-16-47zm-81 117q-34 0-55-15t-21-34q0-13 8.5-21.5T320 202q20 0 20 19q0 3-1 7t-1 6q0 11 14 11q22 0 22-32q0-25-17-25q-13 0-15 9l-44-5l16-89h117l-2 45h-78l-4 23q9-14 41-14q26 0 41 16.5t15 37.5q0 30-21 50t-58 20zm-76-59V96h-85v38h-28v75h15v13h-15v74h128v-74h-15zm-65-106h45v29h-45v-29zm60 160h-88v-34h15v-53h-15v-34h73v87h15v34zm-88 0v-34h15v-20h-16v-30q0-25-16-42.5T140 132q-13 0-25 5V83H2v75h15v64H2v74h209v-20h-15zm-5 0h-80v-34h6v-35q0-20-6-20q-10 0-16 23v32h6v34H22v-34h15V138H22v-35h73v78q15-29 45-29q15 0 25 11.5t10 28.5v50h16v34z"/>`),
		g.Group(children),
	)
}