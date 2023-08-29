package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KanjiUtage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1452 1282h340v-179H761q30-56 45-88h708V411H264v604h264q-14 32-43 88H0v179h384l-55 85q-11 19-11 37q0 26 20 46t38 23q129 13 332 52q-340 76-676 76l49 191q629-23 912-149l36-16q488 122 568 163l138-215q-117-39-429-114q91-88 146-179zm-445 134q-246-53-384-79l33-55h499q-66 83-148 134zm252-647v99H514v-99h745zM514 655v-98h745v98H514zm1259-107V145h-758V0H754v145H6v420h217V325h1341v223h209z"/>`),
		g.Group(children),
	)
}