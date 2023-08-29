package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gdgt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M1 25v276l23 23h253v69H24v46l23 23h276l23-23V324l69-138V48l-69 158V25L323 2H24zm69 230V71h207v184H70z"/>`),
		g.Group(children),
	)
}