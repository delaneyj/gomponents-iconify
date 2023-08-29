package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taskfreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 513q0 72-25 141l76 67q-21 49-54 93l-95-34q-47 56-110 93l16 100q-49 24-101 37l-52-87q-38 6-72 6t-72-6l-52 87q-52-13-101-37l16-100q-63-37-110-93l-95 34q-33-44-54-93l77-67q-26-69-26-141L0 478q4-55 18-107l102-2q26-69 72-124l-50-88q39-40 82-69l80 64q63-36 134-49L457 3q28-3 54.5-3T565 3l19 100q72 13 134 49l80-64q43 29 82 69l-50 88q46 55 72 124l102 2q14 51 18 107zM511 385q-53 0-90.5 37.5T383 513t37.5 90.5T511 641t90.5-37.5T639 513t-37.5-90.5T511 385z"/>`),
		g.Group(children),
	)
}