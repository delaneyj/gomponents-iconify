package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sdvideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 768H128q-53 0-90.5-37.5T0 640V128q0-53 37.5-90.5T128 0h768q53 0 90.5 37.5T1024 128v512q0 53-37.5 90.5T896 768zM224 256h220q-13-55-66-91.5T256 128q-80 0-136 47T64 288t56 113t136 47h32q13 0 22.5 9.5T320 480t-9.5 22.5T288 512H68q13 55 66 91.5T256 640q80 0 136-47t56-113t-56-113t-136-47h-32q-13 0-22.5-9.5T192 288t9.5-22.5T224 256zm736 64q0-80-56-136t-136-56H608q-13 0-22.5 9.5T576 160v448q0 13 9.5 22.5T608 640h160q80 0 136-56t56-136V320zM768 512h-64V256h64q26 0 45 18.5t19 45.5v128q0 27-18.5 45.5T768 512z"/>`),
		g.Group(children),
	)
}