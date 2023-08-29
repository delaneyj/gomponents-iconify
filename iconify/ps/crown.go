package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 475h384q17 0 30-13t13-30V5h-22q-35 0-60 25.5T368 91q0 59 58 81l-94 79l-71-122q18-6 30.5-22.5T304 69q0-27-18.5-45.5T240 5t-45.5 18.5T176 69q0 21 12 37.5t31 22.5l-69 122l-96-79q58-22 58-81q0-35-25-60.5T27 5H5v427q0 17 13 30t30 13zM411 91q0-26 21-37v75q-21-13-21-38zM240 48q21 0 21 21q0 9-6 15.5T240 91q-8 0-14.5-6.5T219 69q0-21 21-21zM48 52q21 11 21 37q0 28-21 38V52zm0 169l113 94l79-139l79 139l113-94v147H48V221zm0 190h384v21H48v-21z"/>`),
		g.Group(children),
	)
}