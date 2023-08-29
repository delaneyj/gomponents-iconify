package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M45 512h342q17 0 29.5-12.5T429 469V192q0-48-31-83.5T321 66q2-4 2-13q0-22-15.5-37.5T269 0H163q-23 0-38.5 15.5T109 53q0 9 2 13q-46 7-77 42.5T3 192v277q0 18 12.5 30.5T45 512zm0-85h43v42H45v-42zm64-43V277h214v192H109v-85zm278 85h-43v-42h43v42zM163 43h106q11 0 11 10q0 11-11 11H163q-11 0-11-11q0-10 11-10zm-32 64h170q35 0 60.5 25t25.5 60v192h-43V203q0-11-11-11q-10 0-10 11v32H109v-32q0-11-10-11q-11 0-11 11v181H45V192q0-35 25.5-60t60.5-25z"/>`),
		g.Group(children),
	)
}