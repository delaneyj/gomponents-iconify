package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emailimport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1017 989L667 640l350-349q7 14 7 29v640q0 15-7 29zM63 256h135q23-111 111-183.5T512 0t203 72.5T826 256h135L512 705zm633 110q8-10 8-23t-8-23H576V160q0-13-9.5-22.5T544 128h-64q-13 0-22.5 9.5T448 160v160H327q-8 10-8 23t8 23l185 146zM7 989q-7-14-7-29V320q0-15 7-29l350 349zm466-233q6 6 15.5 9t16.5 3h7q26 1 39-12l71-71l339 339H63l339-339z"/>`),
		g.Group(children),
	)
}