package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundLevelTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M5 237q0 35 26.5 60.5T95 323h43q14 0 21 10l43 66q18 30 57 30h43q18 0 31.5-12.5T347 387V45q0-17-13-29.5T302 3h-43q-39 0-57 30l-43 66q-7 10-21 10H95q-37 0-63.5 25.5T5 195v42zm43-42q0-18 13.5-30.5T95 152h43q36 0 57-30l43-66q7-11 21-11h45v342h-45q-14 0-21-11l-43-66q-16-30-57-30H95q-20 0-33.5-12.5T48 237v-42zm363 85V152q0-21-22-21q-21 0-21 21v128q0 21 21 21q22 0 22-21zm42-213q-21 0-21 21v256q0 21 21 21q22 0 22-21V88q0-21-22-21z"/>`),
		g.Group(children),
	)
}