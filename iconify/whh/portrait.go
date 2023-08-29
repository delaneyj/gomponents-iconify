package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Portrait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 847V177l128-128v926zM177 128L49 0h926L847 128H177zm-49 719L0 975V49l128 128v670zm704-69q0 10-8.5 18t-26 13.5t-36 9.5t-47 6.5t-50 4t-54.5 2t-50.5.5h-95l-50.5-.5l-54.5-2l-50-4l-47-6.5l-36-9.5l-26-13.5l-8.5-18q2-71 72.5-114T448 611v-13q-67-22-97.5-90T320 343q0-74 53.5-112.5T512 192t138.5 38.5T704 343q0 217-128 256v12q113 10 183.5 53T832 778zm15 118l128 128H49l128-128h670z"/>`),
		g.Group(children),
	)
}