package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.925 7.828a7.862 7.862 0 0 1 1.97 5.215A7.898 7.898 0 0 1 12 20.938a7.899 7.899 0 0 1-7.896-7.895c0-4.189 3.271-7.621 7.396-7.879V4.061H9.913c-.645 0-.643-1 0-1h4.174c.645 0 .644 1 0 1H12.5v1.103a7.865 7.865 0 0 1 4.718 1.956l1.135-1.134a.509.509 0 0 1 .707 0c.199.183.185.522 0 .707l-1.135 1.135Zm.97 5.215A6.898 6.898 0 0 0 12 6.148a6.9 6.9 0 0 0-6.896 6.895A6.898 6.898 0 0 0 12 19.938a6.898 6.898 0 0 0 6.895-6.895Zm-6.395.001c0 .645-1 .643-1 0V8.34c0-.644 1-.643 1 0v4.704Z"/>`),
		g.Group(children),
	)
}