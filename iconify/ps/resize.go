package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m44 514l135-135v56q0 21 21 21t21-21V328q0-21-21-21H93q-21 0-21 21t21 21h56L14 484q-14 14 0 30q8 6 15 6q9 0 15-6zm335-165h56q21 0 21-21t-21-21H328q-21 0-21 21v107q0 21 21 21t21-21v-56l135 135q6 6 15 6q7 0 15-6q14-16 0-30zM484 14L349 149V93q0-21-21-21t-21 21v107q0 21 21 21h107q21 0 21-21t-21-21h-56L514 44q14-14 0-30q-16-14-30 0zM14 14Q0 30 14 44l135 135H93q-21 0-21 21t21 21h107q21 0 21-21V93q0-21-21-21t-21 21v56L44 14Q30 0 14 14z"/>`),
		g.Group(children),
	)
}