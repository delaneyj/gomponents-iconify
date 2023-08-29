package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailDownloadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12.803A6 6 0 0 0 13.803 21H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v8.803Zm-9.94-1.12L5.648 6.238L4.353 7.762l7.72 6.555l7.581-6.56l-1.308-1.513l-6.285 5.439ZM20 18h3l-4 4l-4-4h3v-4h2v4Z"/>`),
		g.Group(children),
	)
}