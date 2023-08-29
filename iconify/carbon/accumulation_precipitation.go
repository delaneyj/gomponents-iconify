package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccumulationPrecipitation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 16h2v2h-2zm2 2h2v2h-2zm2 2h2v2h-2zm0-4h2v2h-2zm-4 4h2v2h-2zm4-8a3.898 3.898 0 0 1-4-3.777a3.902 3.902 0 0 1 .653-2.064l2.517-3.745a1.038 1.038 0 0 1 1.66 0l2.485 3.696A3.97 3.97 0 0 1 24 8.223A3.898 3.898 0 0 1 20 12zm0-7.237l-1.656 2.463a1.89 1.89 0 0 0-.344.997A1.9 1.9 0 0 0 20 10a1.9 1.9 0 0 0 2-1.777a1.98 1.98 0 0 0-.375-1.047z"/><path fill="currentColor" d="M28 4a2.002 2.002 0 0 0-2 2v20H6V10h2v8l1 2l1-2v-8h2v4l1 2l1-2V8H6V6a2.002 2.002 0 0 0-2-2H2v2h2v20a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V6h2V4Z"/>`),
		g.Group(children),
	)
}