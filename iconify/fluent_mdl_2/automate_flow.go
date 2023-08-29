package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutomateFlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M821 768h568L313 1792H56l304-640H0L535 128h612L821 768zm-559 896l807-768H613l325-640H613l-402 768h351l-304 640h4zM2048 640h-256v640h256v640h-640v-640h256v-256h-384l128-128h256V640h-256V0h640v640zm-128 768h-384v384h384v-384zm-384-896h384V128h-384v384z"/>`),
		g.Group(children),
	)
}