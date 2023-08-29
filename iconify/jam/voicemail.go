package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.743 9h4.514A5.5 5.5 0 1 1 19 10.978V11H6v-.022A5.5 5.5 0 1 1 9.743 9zM5.5 9a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7zm13 0a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7z"/>`),
		g.Group(children),
	)
}