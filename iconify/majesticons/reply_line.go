package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10.383 4.076A1 1 0 0 1 11 5v3.028c4.117.23 6.861 1.885 8.58 4.035C21.403 14.339 22 17.075 22 19a1 1 0 1 1-2 0c0-.502-.235-.942-.802-1.366c-.592-.443-1.477-.805-2.565-1.076c-1.804-.45-3.941-.594-5.633-.583V19a1 1 0 0 1-1.707.707l-7-7a1 1 0 0 1 0-1.414l7-7a1 1 0 0 1 1.09-.217zm8.843 11.267a8.762 8.762 0 0 0-1.207-2.03C16.574 11.505 14.122 10 10 10a1 1 0 0 1-1-1V7.414L4.414 12L9 16.586V15a1 1 0 0 1 .955-.999c1.888-.086 4.743.014 7.162.616a11.6 11.6 0 0 1 2.11.726z"/></g>`),
		g.Group(children),
	)
}