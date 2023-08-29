package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCostaRica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M0 22h64v20H0z"/><path fill="#2b3990" d="M10 54h44c5.086 0 8.247-2.905 9.446-7H.554c1.199 4.095 4.36 7 9.446 7m44-44H10c-5.086 0-8.247 2.905-9.446 7h62.893c-1.2-4.095-4.361-7-9.447-7"/><path fill="#e6e7e8" d="M64 42H0v1c0 1.413.19 2.759.554 4h62.893c.363-1.241.553-2.587.553-4v-1m0-21c0-1.413-.19-2.759-.554-4H.554A14.215 14.215 0 0 0 0 21v1h64v-1z"/>`),
		g.Group(children),
	)
}