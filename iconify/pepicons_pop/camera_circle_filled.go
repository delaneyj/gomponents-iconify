package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCameraCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.5 7h-7a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3Zm-8 3a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-6Z"/><path d="m19.934 8.176l-3.468 2.381a1 1 0 0 0-.434.815L16 14.587a1 1 0 0 0 .434.834l3.5 2.403A1 1 0 0 0 21.5 17V9a1 1 0 0 0-1.566-.824ZM19.5 15.1l-1.495-1.026l.022-2.163L19.5 10.9v4.2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCameraCircleFilled0)"/></g>`),
		g.Group(children),
	)
}