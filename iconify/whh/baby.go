package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M751.5 622.5Q734 640 709.5 640T667 622L513 467v173H257V467L102 622q-18 18-42.5 18t-42-17.5T0 580t17-42l144-154q64-64 224-64t224 64l143 154q17 17 17 42t-17.5 42.5zM384.5 256q-53.5 0-91-37.5T256 128t37.5-90.5t91-37.5t91 37.5T513 128t-37.5 90.5t-91 37.5zM82 785q3-2 13-8l118-100l125 90l-112 74l76 77q18 18 18 43.5t-18 44t-44 18.5t-44-18L82 874q-18-19-18-44.5T82 785zm385 133l76-77l-112-74l125-90l118 100q10 6 13 8q18 19 18 44.5T687 874l-132 132q-18 18-44 18t-44-18t-18-44t18-44z"/>`),
		g.Group(children),
	)
}