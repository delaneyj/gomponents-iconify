package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.477 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10zM21 39h-2.553l-5.084-9.131V39H11V25h2.477l5.16 9.348V25H21v14zm14 0H25V25h9.75v2.367h-7.096v3.104h6.6v2.359h-6.6v3.808H35V39zm16.668 0H48.73l-2.729-10.467L43.279 39h-3.004L37 25h2.836l2.068 9.615L44.41 25h3.293l2.404 9.779L52.213 25H55l-3.332 14z"/>`),
		g.Group(children),
	)
}