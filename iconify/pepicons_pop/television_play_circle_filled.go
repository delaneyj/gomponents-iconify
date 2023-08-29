package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionPlayCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTelevisionPlayCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M7 11v7a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1Zm-2 7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v7Z"/><path d="M12.875 8.78a1 1 0 0 1-.156-1.405l2-2.5a1 1 0 0 1 1.562 1.25l-2 2.5a1 1 0 0 1-1.406.156Z"/><path d="M13.125 8.78a1 1 0 0 0 .156-1.405l-2-2.5a1 1 0 1 0-1.562 1.25l2 2.5a1 1 0 0 0 1.406.156ZM15 14.5l-3.25 2v-4l3.25 2Z"/><path d="M15.5 14.5a.5.5 0 0 1-.238.426l-3.25 2a.5.5 0 0 1-.762-.426v-4a.5.5 0 0 1 .762-.426l3.25 2a.5.5 0 0 1 .238.426Zm-3.25-1.105v2.21l1.796-1.105l-1.796-1.105Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTelevisionPlayCircleFilled0)"/></g>`),
		g.Group(children),
	)
}