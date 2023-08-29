package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 2.125l-.906 2.063l-3.25 7.28l-7.938.845l-2.25.25l1.688 1.5l5.906 5.343l-1.656 7.813l-.469 2.187l1.969-1.125l6.906-4l6.906 4l1.969 1.125l-.469-2.187l-1.656-7.813l5.906-5.343l1.688-1.5l-2.25-.25l-7.938-.844l-3.25-7.281zm0 4.906l2.563 5.782l.25.53l.562.063l6.281.656l-4.687 4.22l-.438.405l.125.563l1.313 6.156l-5.469-3.125l-.5-.312l-.5.312l-5.469 3.125l1.313-6.156l.125-.563l-.438-.406l-4.687-4.218l6.281-.657l.563-.062l.25-.531z"/>`),
		g.Group(children),
	)
}