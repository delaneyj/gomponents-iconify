package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePDB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1920h896v128H128V0h1115l549 549v475h-128V640h-512V128H256v1792zM1280 512h293l-293-293v293zm320 640q93 0 174 35t143 96t96 142t35 175q0 93-35 174t-96 143t-142 96t-175 35q-93 0-174-35t-143-96t-96-142t-35-175q0-93 35-174t96-143t142-96t175-35z"/>`),
		g.Group(children),
	)
}