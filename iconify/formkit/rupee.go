package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rupee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7ZM8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6Z"/><path fill="currentColor" d="M6.5 9h-1c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h1C7.33 8 8 7.33 8 6.5S7.33 5 6.5 5h-1c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h1a2.5 2.5 0 0 1 0 5Z"/><path fill="currentColor" d="M10.5 7h-5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5Zm0-2h-5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5Zm-3 7c-.16 0-.32-.08-.42-.22l-2-3a.498.498 0 0 1 .83-.55l2 3c.15.23.09.54-.14.69c-.09.06-.18.08-.28.08Z"/>`),
		g.Group(children),
	)
}