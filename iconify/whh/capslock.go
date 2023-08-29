package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capslock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M869 513H703v191q0 27-18.5 45.5T640 768H256q-26 0-45-18.5T192 704V513H28q-12 0-20-13t-7-31t13-30L403 19q18-19 45-19t46 19l388 420q12 12 13.5 30t-6.5 31t-20 13zM256 831h384q26 0 44.5 19t18.5 45v65q0 26-18.5 45t-44.5 19H256q-26 0-45-18.5T192 960v-65q0-26 19-45t45-19z"/>`),
		g.Group(children),
	)
}