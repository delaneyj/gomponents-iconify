package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7ZM8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6Z"/><path fill="currentColor" d="M9.62 11H6.37c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.25c.28 0 .5.22.5.5s-.22.5-.5.5Zm0-2H6.37c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.25c.28 0 .5.22.5.5s-.22.5-.5.5Z"/><path fill="currentColor" d="M8 12.5c-.28 0-.5-.22-.5-.5V8.5c0-.28.22-.5.5-.5s.5.22.5.5V12c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M8 9a.42.42 0 0 1-.26-.08a.501.501 0 0 1-.16-.69l2.5-4a.501.501 0 0 1 .85.53l-2.5 4c-.09.15-.26.24-.42.24Z"/><path fill="currentColor" d="M8 9a.47.47 0 0 1-.42-.24l-2.5-4a.495.495 0 0 1 .16-.69c.23-.15.54-.08.69.16l2.5 4c.15.23.08.54-.16.69c-.08.05-.17.08-.26.08Z"/>`),
		g.Group(children),
	)
}