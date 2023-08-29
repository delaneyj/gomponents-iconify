package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssFeedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19q0-.825.588-1.413T5 17q.825 0 1.413.588T7 19q0 .825-.588 1.413T5 21Zm13.475 0q-.6 0-1.038-.425t-.512-1.05q-.25-2.45-1.313-4.6t-2.724-3.813Q11.225 9.45 9.075 8.389t-4.6-1.313Q3.85 7 3.425 6.55T3 5.475q0-.625.438-1.038t1.037-.362q3.075.25 5.775 1.55t4.763 3.363q2.062 2.062 3.362 4.762t1.55 5.775q.05.6-.375 1.038T18.475 21Zm-6 0q-.575 0-1.037-.413t-.588-1.062q-.45-2.425-2.2-4.175t-4.175-2.2q-.65-.125-1.062-.6T3 11.475q0-.65.45-1.05t1.025-.325q3.7.5 6.313 3.113t3.112 6.312q.075.575-.35 1.025t-1.075.45Z"/>`),
		g.Group(children),
	)
}