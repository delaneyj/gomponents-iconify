package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comparison(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-.25 2.386a2.501 2.501 0 1 0-1.5 0V10.5a2.75 2.75 0 0 0 2.75 2.75h.629l-.47.47a.75.75 0 1 0 1.061 1.06l1.75-1.75l.53-.53l-.53-.53l-1.75-1.75a.747.747 0 0 0-1.061 0a.75.75 0 0 0 0 1.06l.47.47H5.5c-.69 0-1.25-.56-1.25-1.25V5.886ZM9.28 1.22a.75.75 0 1 1 1.06 1.06l-.47.47h.63a2.75 2.75 0 0 1 2.75 2.75v4.614a2.501 2.501 0 1 1-1.5 0V5.5c0-.69-.56-1.25-1.25-1.25h-.63l.47.47a.75.75 0 1 1-1.06 1.06L7.53 4.03L7 3.5l.53-.53l1.75-1.75ZM12.5 13.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}