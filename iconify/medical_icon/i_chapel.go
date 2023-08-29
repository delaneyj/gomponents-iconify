package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IChapel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M39.65 35.391V16.028L32.203.482l-7.439 15.546v19.363L9.66 42.686v20.07h17.916v-9.678c0-4.963 4.492-6.588 4.492-6.588s4.491 1.625 4.491 6.588v9.678h18.202v-20.07l-15.11-7.295zm-4.905-7.908h-5.103v-4.097c0-4.109 2.551-5.456 2.551-5.456s2.552 1.347 2.552 5.456v4.097z"/>`),
		g.Group(children),
	)
}