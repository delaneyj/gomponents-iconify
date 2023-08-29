package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func User(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g stroke="currentColor"><path fill="none" stroke-linecap="round" stroke-miterlimit="2.3" stroke-width="70" d="M722.297 656.594C722.317 737.279 768 840 821.517 880.15c14.362 10.775 1.471 7.63 18.999 18.584C704.714 939.974 528 1010 448 1160v420h1087.5v-420c-80-150-256.714-220.026-392.515-261.266c96-60 118.195-143.433 118.218-242.14C1260.89 520 1151.5 400 991.75 400C832 400 722.61 520 722.297 656.594z"/><path fill="currentColor" fill-rule="evenodd" d="M752 550c38.64 3.52 39.962-25.563 145.847-19.322c158.153 9.322 107.028 57.135 214.93 60.28c84.098-3.145 80.524-29.91 119.223-40.958c-64-70-127.297-114.535-240-150c-96 30-176 80-240 150z"/><path fill="none" stroke-width="70" d="M684 1560v-280m614 280v-280M768 920s-37.712 122.288 0 160c105.595 105.595 342.405 105.595 448 0c37.712-37.712 0-160 0-160"/></g>`),
		g.Group(children),
	)
}