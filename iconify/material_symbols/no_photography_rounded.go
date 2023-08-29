package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoPhotographyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.575L18.175 21H4q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h1.025v2.85L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM22 19.125l-5.525-5.55q.125-1.05-.213-2.037T15.176 9.8q-.725-.725-1.713-1.063t-2.037-.212L7.5 4.625l.9-.975q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12.125ZM12 17.5q.575 0 1.113-.125t1.037-.4L8.025 10.85q-.275.5-.4 1.038T7.5 13q0 1.875 1.313 3.188T12 17.5Zm0-2q-1.05 0-1.775-.725T9.5 13q0-.5.188-.962t.537-.813l3.55 3.55q-.35.35-.812.538T12 15.5Z"/>`),
		g.Group(children),
	)
}