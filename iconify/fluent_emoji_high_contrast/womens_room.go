package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomensRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15.941 10a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-4.808 2.424a2 2 0 0 1 1.933-1.486h5.904a2 2 0 0 1 1.93 1.475l1.159 4.262a1 1 0 0 1-1.93.525l-1.056-3.885a.461.461 0 0 0-.495-.335a.455.455 0 0 0-.406.45h-.001l1.892 6.954a.5.5 0 0 1-.483.632H18.5V26a1 1 0 0 1-1.048.999c-.538-.025-.952-.488-.952-1.027v-4.456a.5.5 0 1 0-1 0V26a1 1 0 0 1-1.048.999c-.538-.025-.952-.488-.952-1.027v-4.956h-1.06a.5.5 0 0 1-.481-.635l1.932-6.912h-.008a.472.472 0 0 0-.445-.471l-.126-.007a.303.303 0 0 0-.31.224l-1.067 3.98a1 1 0 1 1-1.933-.514l1.131-4.257Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}