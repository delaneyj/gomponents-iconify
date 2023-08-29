package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IScheduleSchoolDateTimeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsIScheduleSchoolDateTimeNegative0)"><path fill-rule="evenodd" d="M35 40a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm1-7.5a1 1 0 1 0-2 0v2.914l1.293 1.293a1 1 0 0 0 1.414-1.414L36 34.586V32.5Z" clip-rule="evenodd"/><path d="M12 23v2h2v-2h-2Zm8 0h2v2h-2v-2Zm8 0v2h2v-2h-2Zm-14 8v2h-2v-2h2Zm8 0h-2v2h2v-2Z"/><path fill-rule="evenodd" d="M8 18v18a1 1 0 0 0 1 1h19.29A7.001 7.001 0 0 1 34 28.07V18H8Zm4 3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-2Zm6 2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2Zm10-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-2ZM10 31a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2Zm10-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm35 42a7 7 0 0 0 1-13.93V12a3 3 0 0 0-3-3h-3v3a1 1 0 1 1-2 0V7a1 1 0 1 0-2 0v2H16v3a1 1 0 1 1-2 0V7a1 1 0 1 0-2 0v2H9a3 3 0 0 0-3 3v24a3 3 0 0 0 3 3h20.255A6.992 6.992 0 0 0 35 42Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsIScheduleSchoolDateTimeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}