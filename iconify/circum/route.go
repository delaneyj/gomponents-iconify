package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Route(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.792 17.086c-.58-.58-1.16-1.17-1.75-1.75c-.08-.08-.16-.17-.25-.25a.492.492 0 0 0-.7 0a.5.5 0 0 0 0 .71l1.14 1.14H9.282a2.22 2.22 0 0 1 0-4.44h3a3.215 3.215 0 1 0 0-6.43h-5.27a2.5 2.5 0 1 0 0 1h5.27a2.215 2.215 0 1 1 0 4.43h-3a3.22 3.22 0 1 0 0 6.44h10.96l-.9.9c-.09.08-.17.17-.25.25a.5.5 0 0 0 0 .71a.511.511 0 0 0 .7 0l1.75-1.75l.25-.25a.5.5 0 0 0 0-.71Zm-17.23-9.02a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Z"/>`),
		g.Group(children),
	)
}