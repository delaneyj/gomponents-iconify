package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationBarriers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 128h1664v640h-128v320l-320-320h-64V640h117l139 139V640h128V256H512v299q-42 47-73 100t-52 113h-3V128zm576 1280q-93 0-174-35t-143-96t-96-142t-35-175q0-93 35-174t96-143t142-96t175-35q93 0 174 35t143 96t96 142t35 175q0 93-35 174t-96 143t-142 96t-175 35zm320-448q0-66-25-124t-69-101t-102-69t-124-26q-47 0-92 13t-84 40l443 443q26-39 39-84t14-92zm-640 0q0 66 25 124t68 102t102 69t125 25q47 0 92-13t84-40L693 784q-26 39-39 84t-14 92zm896 576v-512h128v640H448l-320 320v-320H0v-640h384v128H128v384h128v139l139-139h1141z"/>`),
		g.Group(children),
	)
}