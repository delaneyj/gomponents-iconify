package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abuse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.308 1.5L1.5 5.308v5.384L5.308 14.5h5.384l3.808-3.808V5.308L10.692 1.5H5.308ZM11.314 0H4.686L0 4.686v6.628L4.686 16h6.628L16 11.314V4.686L11.314 0ZM9 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-.25-6.25a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0v-3.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}