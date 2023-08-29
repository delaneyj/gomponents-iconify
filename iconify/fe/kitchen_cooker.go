package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenCooker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feKitchenCooker0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feKitchenCooker1" fill="currentColor" fill-rule="nonzero"><path id="feKitchenCooker2" d="M4.268 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-.268a2 2 0 0 1-3.464 0H7.732a2 2 0 0 1-3.464 0ZM4 6v12h16V6H4Zm2 6h12v4H6v-4Zm0-4h5v2H6V8Zm4 5v1h4v-1h-4Zm4-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/></g></g>`),
		g.Group(children),
	)
}