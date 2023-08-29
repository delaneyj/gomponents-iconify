package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireHose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.27 10.568a2.31 2.31 0 0 1-2.308 2.312a2.31 2.31 0 0 1-2.308-2.312A3.85 3.85 0 0 1 11.5 6.714a3.85 3.85 0 0 1 3.846 3.854a5.39 5.39 0 0 1-5.384 5.396a5.39 5.39 0 0 1-5.385-5.396c0-3.832 3.1-6.938 6.923-6.938s6.923 3.106 6.923 6.938c0 4.683-3.788 8.479-8.461 8.479c-4.674 0-8.462-3.796-8.462-8.48c0-5.534 4.477-10.02 10-10.02s10 4.486 10 10.02v10.35l-.333.347c-.658.684-.985 1.65-1.167 2.6c-.183-.95-.509-1.916-1.167-2.6l-.333-.347v-4.87m0 2.5h3"/>`),
		g.Group(children),
	)
}