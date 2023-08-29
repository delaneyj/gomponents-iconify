package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 13.5v-1.879l7.28-7.28l1.88 1.879l-7.28 7.28H2.5Zm10.22-8.341l.805-.805a.5.5 0 0 0 0-.708l-1.171-1.171a.5.5 0 0 0-.708 0l-.805.805l1.879 1.88ZM1 13.5V11l9.586-9.586a2 2 0 0 1 2.828 0l1.172 1.172a2 2 0 0 1 0 2.828L5 15H1v-1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}