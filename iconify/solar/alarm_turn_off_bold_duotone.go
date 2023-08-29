package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmTurnOffBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z" opacity=".5"/><path d="M14.652 10.349a.75.75 0 0 1 0 1.06L13.06 13l1.59 1.591a.75.75 0 0 1-1.06 1.06L12 14.061l-1.592 1.59a.75.75 0 0 1-1.06-1.06l1.59-1.59l-1.59-1.592a.75.75 0 0 1 1.06-1.06L12 11.939l1.591-1.59a.75.75 0 0 1 1.06 0Z"/><path fill-rule="evenodd" d="M8.24 2.34a.719.719 0 0 1-.232.996l-3.891 2.41a.734.734 0 0 1-1.006-.23a.719.719 0 0 1 .232-.996l3.892-2.41a.734.734 0 0 1 1.006.23Zm7.52 0a.734.734 0 0 1 1.005-.23l3.892 2.41a.719.719 0 0 1 .232.996a.734.734 0 0 1-1.006.23l-3.891-2.41a.719.719 0 0 1-.233-.996Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}