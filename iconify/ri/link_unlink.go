package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkUnlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 17h5v2h-3v3h-2v-5ZM7 7H2V5h3V2h2v5Zm11.364 8.536L16.95 14.12l1.414-1.414a5 5 0 0 0-7.071-7.071L9.879 7.05L8.464 5.636L9.88 4.222a7 7 0 0 1 9.9 9.9l-1.415 1.414Zm-2.829 2.828l-1.414 1.414a7 7 0 0 1-9.9-9.9l1.415-1.414L7.05 9.88l-1.414 1.414a5 5 0 0 0 7.071 7.071l1.414-1.414l1.415 1.414Zm-.707-10.607l1.415 1.415l-7.071 7.07l-1.415-1.414l7.071-7.07Z"/>`),
		g.Group(children),
	)
}