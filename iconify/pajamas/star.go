package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.454 1.694a.591.591 0 0 1 1.092 0l1.585 3.81a.25.25 0 0 0 .21.154l4.114.33a.591.591 0 0 1 .338 1.038L11.658 9.71a.25.25 0 0 0-.08.247l.957 4.015a.591.591 0 0 1-.883.641l-3.522-2.15a.25.25 0 0 0-.26 0l-3.522 2.15a.591.591 0 0 1-.883-.641l.957-4.015a.25.25 0 0 0-.08-.247L1.207 7.026a.591.591 0 0 1 .338-1.038l4.113-.33a.25.25 0 0 0 .211-.153l1.585-3.81Z"/>`),
		g.Group(children),
	)
}