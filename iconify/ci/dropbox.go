package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.006 20.5L7 17.311l5-3.186l5 3.186l-5 3.188l.006.001Zm5-4.251l-5-3.187l5-3.187l-5-3.186l5-3.189l5 3.189l-5 3.186l5 3.187l-5 3.187ZM7 16.249l-5-3.187l5-3.187l-5-3.186L7 3.5l5 3.189l-5 3.185l5 3.187l-5 3.188Z"/>`),
		g.Group(children),
	)
}