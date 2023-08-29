package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SupervisorAccountOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 15q-1.05 0-1.775-.725T14.5 12.5q0-1.05.725-1.775T17 10q1.05 0 1.775.725T19.5 12.5q0 1.05-.725 1.775T17 15Zm-5 5v-1.4q0-.6.313-1.113t.887-.737q.9-.375 1.863-.563T17 16q.975 0 1.938.188t1.862.562q.575.225.888.738T22 18.6V20H12Zm-2-8q-1.65 0-2.825-1.175T6 8q0-1.65 1.175-2.825T10 4q1.65 0 2.825 1.175T14 8q0 1.65-1.175 2.825T10 12Zm0-4ZM2 20v-2.8q0-.85.425-1.563T3.6 14.55q1.5-.75 3.112-1.15T10 13q.875 0 1.75.15t1.75.35l-.85.85l-.85.85q-.45-.125-.9-.163T10 15q-1.45 0-2.838.35t-2.662 1q-.25.125-.375.35T4 17.2v.8h6v2H2Zm8-2Zm0-8q.825 0 1.413-.588T12 8q0-.825-.588-1.413T10 6q-.825 0-1.413.588T8 8q0 .825.588 1.413T10 10Z"/>`),
		g.Group(children),
	)
}