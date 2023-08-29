package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 0v128h128L320 0zm-64 405.3c35.4 0 64-28.6 64-64c0-9.4-2.1-18.2-5.7-26.2l-84.5 84.5c8 3.6 16.8 5.7 26.2 5.7zM298.7 0H64v512h384V149.3H298.7V0zm64 341.3c0 58.9-47.8 106.7-106.7 106.7c-58.9 0-106.7-47.7-106.7-106.7c0-58.9 47.8-106.7 106.7-106.7c58.9.1 106.7 47.8 106.7 106.7zm-170.7 0c0 9.4 2.1 18.2 5.7 26.2l84.5-84.5c-8-3.6-16.9-5.8-26.2-5.8c-35.4.1-64 28.8-64 64.1z"/>`),
		g.Group(children),
	)
}