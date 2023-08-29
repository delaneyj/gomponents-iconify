package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gameboy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M544 1024H64q-26 0-45-19T0 960V64q0-27 18.5-45.5T64 0h640q27 0 45.5 18.5T768 64v736q0 93-65.5 158.5T544 1024zm-31.5-128q26.5 0 45-19t18.5-45.5t-18.5-45t-45-18.5t-45.5 18.5t-19 45t19 45.5t45.5 19zM64 720v32q0 7 4.5 11.5T80 768h48v48q0 7 4.5 11.5T144 832h32q7 0 11.5-4.5T192 816v-48h48q7 0 11.5-4.5T256 752v-32q0-7-4.5-11.5T240 704h-48v-48q0-7-4.5-11.5T176 640h-32q-7 0-11.5 4.5T128 656v48H80q-7 0-11.5 4.5T64 720zm576-560q0-13-9.5-22.5T608 128H160q-13 0-22.5 9.5T128 160v320q0 13 9.5 22.5T160 512h448q13 0 22.5-9.5T640 480V160zm.5 480q-26.5 0-45.5 18.5t-19 45t19 45.5t45.5 19t45-19t18.5-45.5t-18.5-45t-45-18.5z"/>`),
		g.Group(children),
	)
}