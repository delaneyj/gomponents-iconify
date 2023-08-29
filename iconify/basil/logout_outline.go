package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoutOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.25a.75.75 0 0 0 0 1.5h6A1.75 1.75 0 0 0 19.75 18V6A1.75 1.75 0 0 0 18 4.25h-6a.75.75 0 0 0 0 1.5h6a.25.25 0 0 1 .25.25v12a.25.25 0 0 1-.25.25h-6Z"/><path fill="currentColor" fill-rule="evenodd" d="M14.503 14.365c.69 0 1.25-.56 1.25-1.25v-2.24c0-.69-.56-1.25-1.25-1.25H9.89a26.723 26.723 0 0 0-.02-.22l-.054-.556a1.227 1.227 0 0 0-1.751-.988a15.059 15.059 0 0 0-4.368 3.163l-.099.104a1.253 1.253 0 0 0 0 1.734l.1.103a15.06 15.06 0 0 0 4.367 3.164a1.227 1.227 0 0 0 1.751-.988l.054-.556l.02-.22h4.613Zm-5.308-1.5a.75.75 0 0 0-.748.704c-.019.29-.042.581-.07.871l-.016.162a13.562 13.562 0 0 1-3.516-2.607a13.558 13.558 0 0 1 3.516-2.607l.016.162c.028.29.051.58.07.871a.75.75 0 0 0 .748.704h5.058v1.74H9.195Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}