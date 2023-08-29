package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psltwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 768H192q-80 0-136-56T0 576V192q0-80 56-136T192 0h640q80 0 136 56t56 136v384q0 80-56 136t-136 56zM416 576H256V160q0-13-9.5-22.5T224 128t-22.5 9.5T192 160v448q0 13 9.5 22.5T224 640h192q13 0 22.5-9.5T448 608t-9.5-22.5T416 576zm288-128q53 0 90.5-37.5T832 320v-64q0-53-37.5-90.5T704 128t-90.5 37.5T576 256q0 13 9.5 22.5T608 288t22.5-9.5T640 256q0-27 18.5-45.5T704 192t45.5 19t18.5 45v64q0 26-18.5 45T704 384q-53 0-90.5 37.5T576 512v96q0 13 9.5 22.5T608 640h192q13 0 22.5-9.5T832 608t-9.5-22.5T800 576H656q-7 0-11.5-4.5T640 560v-48q0-27 18.5-45.5T704 448z"/>`),
		g.Group(children),
	)
}