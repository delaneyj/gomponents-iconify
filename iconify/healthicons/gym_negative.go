package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GymNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsGymNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm23 7h3v16h-3V16Zm8 3h-3v10h3v-4h1v-2h-1v-4ZM19 32h-3V16h3v16Zm-8-3h3V19h-3v4h-1v2h1v4Zm16-4h-6v-2h6v2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsGymNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}