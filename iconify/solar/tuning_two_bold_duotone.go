package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9.25 14a3 3 0 1 1 0 6a3 3 0 0 1 0-6Zm5-10a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/><path d="M17.166 7.708a3.008 3.008 0 0 0-.021-1.5h4.605a.75.75 0 0 1 0 1.5h-4.584Zm-5.811-1.5a3.003 3.003 0 0 0-.02 1.5H1.75a.75.75 0 0 1 0-1.5h9.605ZM6.356 16.209H1.75a.75.75 0 0 0 0 1.5h4.584a3.007 3.007 0 0 1 .022-1.5Zm5.81 1.5h9.584a.75.75 0 0 0 0-1.5h-9.605a3.003 3.003 0 0 1 .02 1.5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}