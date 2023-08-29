package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quota(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.457 8.75a6.501 6.501 0 0 1-12.914 0h.707a.75.75 0 0 0 0-1.5h-.707a6.469 6.469 0 0 1 1.361-3.285l.5.5a.75.75 0 0 0 1.06-1.061l-.5-.5A6.469 6.469 0 0 1 7.25 1.543v.707a.75.75 0 0 0 1.5 0v-.707a6.47 6.47 0 0 1 3.285 1.361l-.5.5a.75.75 0 1 0 1.061 1.06l.5-.5a6.469 6.469 0 0 1 1.361 3.286h-.707a.75.75 0 0 0 0 1.5h.707ZM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0Zm-8 4a2 2 0 0 0 .75-3.855V4.75a.75.75 0 0 0-1.5 0v3.395A2 2 0 0 0 8 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}