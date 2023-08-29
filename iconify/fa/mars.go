package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1472 0q26 0 45 19t19 45v416q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-382 383q126 156 126 359q0 117-45.5 223.5t-123 184t-184 123T576 1536t-223.5-45.5t-184-123t-123-184T0 960t45.5-223.5t123-184t184-123T576 384q203 0 359 126l382-382h-261q-14 0-23-9t-9-23V32q0-14 9-23t23-9h416zM576 1408q185 0 316.5-131.5T1024 960T892.5 643.5T576 512T259.5 643.5T128 960t131.5 316.5T576 1408z"/>`),
		g.Group(children),
	)
}