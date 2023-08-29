package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Datetime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.5 14h-8C.67 14 0 13.33 0 12.5V2.38C0 1.55.67.88 1.5.88h11c.83 0 1.5.67 1.5 1.5v7.25c0 .28-.22.5-.5.5s-.5-.22-.5-.5V2.38c0-.28-.22-.5-.5-.5h-11c-.28 0-.5.22-.5.5V12.5c0 .28.22.5.5.5h8c.28 0 .5.22.5.5s-.22.5-.5.5Z"/><path fill="currentColor" d="M4 3.62c-.28 0-.5-.22-.5-.5V.5c0-.28.22-.5.5-.5s.5.22.5.5v2.62c0 .28-.22.5-.5.5Zm6.12 0c-.28 0-.5-.22-.5-.5V.5c0-.28.22-.5.5-.5s.5.22.5.5v2.62c0 .28-.22.5-.5.5ZM13.5 6H.5C.22 6 0 5.78 0 5.5S.22 5 .5 5h13c.28 0 .5.22.5.5s-.22.5-.5.5Zm-1 10C10.57 16 9 14.43 9 12.5S10.57 9 12.5 9s3.5 1.57 3.5 3.5s-1.57 3.5-3.5 3.5Zm0-6a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5Z"/><path fill="currentColor" d="M13.5 14a.47.47 0 0 1-.35-.15l-1-1a.51.51 0 0 1-.15-.35V11c0-.28.22-.5.5-.5s.5.22.5.5v1.29l.85.85c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		g.Group(children),
	)
}