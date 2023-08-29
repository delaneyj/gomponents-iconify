package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmAddBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z" opacity=".5"/><path d="M12 9.25a.75.75 0 0 1 .75.75v2.25H15a.75.75 0 0 1 0 1.5h-2.25V16a.75.75 0 0 1-1.5 0v-2.25H9a.75.75 0 0 1 0-1.5h2.25V10a.75.75 0 0 1 .75-.75Z"/><path fill-rule="evenodd" d="M8.24 2.34a.719.719 0 0 1-.232.996l-3.891 2.41a.734.734 0 0 1-1.006-.23a.719.719 0 0 1 .232-.996l3.892-2.41a.734.734 0 0 1 1.006.23Zm7.52 0a.734.734 0 0 1 1.005-.23l3.892 2.41a.719.719 0 0 1 .232.996a.734.734 0 0 1-1.006.23l-3.891-2.41a.719.719 0 0 1-.233-.996Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}