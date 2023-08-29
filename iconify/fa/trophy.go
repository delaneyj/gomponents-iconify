package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M458 755q-74-162-74-371H128v96q0 78 94.5 162T458 755zm1078-275v-96h-256q0 209-74 371q141-29 235.5-113t94.5-162zm128-128v128q0 71-41.5 143t-112 130t-173 97.5T1122 895q-42 54-95 95q-38 34-52.5 72.5T960 1152q0 54 30.5 91t97.5 37q75 0 133.5 45.5T1280 1440v64q0 14-9 23t-23 9H416q-14 0-23-9t-9-23v-64q0-69 58.5-114.5T576 1280q67 0 97.5-37t30.5-91q0-51-14.5-89.5T637 990q-53-41-95-95q-113-5-215.5-44.5t-173-97.5t-112-130T0 480V352q0-40 28-68t68-28h288v-96q0-66 47-113T544 0h576q66 0 113 47t47 113v96h288q40 0 68 28t28 68z"/>`),
		g.Group(children),
	)
}