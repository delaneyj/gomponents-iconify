package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReleaseGateError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m749 403l557 557l-557 557l-90-90l402-403H0V896h1061L659 493l90-90zm787 621V128h-384V0h512v1024h-128zm512 576q0 93-35 174t-96 143t-142 96t-175 35q-93 0-174-35t-143-96t-96-142t-35-175q0-93 35-174t96-143t142-96t175-35q93 0 174 35t143 96t96 142t35 175zm-448 320q66 0 124-25t101-68t69-102t26-125q0-66-25-124t-69-101t-102-69t-124-26q-66 0-124 25t-102 69t-69 102t-25 124q0 66 25 124t68 102t102 69t125 25zm104-504l80 80l-105 104l105 104l-80 80l-104-105l-104 105l-80-80l105-104l-105-104l80-80l104 105l104-105z"/>`),
		g.Group(children),
	)
}