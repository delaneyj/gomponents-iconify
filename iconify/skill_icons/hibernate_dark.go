package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HibernateDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path fill="#59666C" d="m75.616 38l34.622 59.962l-34.639 60.047L41 97.962L75.616 38Z"/><path fill="#BCAE79" d="M144.871 38.003H75.636l34.622 59.962h69.254l-34.641-59.962Z"/><path fill="#59666C" d="m179.476 218l-34.622-59.962l34.639-60.047l34.599 60.047L179.476 218Z"/><path fill="#BCAE79" d="M110.221 217.995h69.244l-34.594-59.962H75.58l34.641 59.962Z"/></g>`),
		g.Group(children),
	)
}