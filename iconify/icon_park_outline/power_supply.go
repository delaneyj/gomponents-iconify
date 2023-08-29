package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSupply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32.297 6H15.704a1 1 0 0 0-.942.662l-6.67 18.581a1 1 0 0 0-.04.525l1.793 9.42a1 1 0 0 0 .982.812h26.346a1 1 0 0 0 .982-.813l1.794-9.419a1 1 0 0 0-.041-.525l-6.67-18.58A1 1 0 0 0 32.297 6ZM24 36v8m0-32v6m-7.5 5.402l3 5.196m12-5.196l-3 5.196"/>`),
		g.Group(children),
	)
}