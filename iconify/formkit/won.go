package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Won(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7ZM8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6Z"/><path fill="currentColor" d="M10 12.5H9c-.25 0-.46-.18-.49-.42L8 8.79l-.51 3.29c-.04.24-.25.42-.49.42H6a.51.51 0 0 1-.49-.39l-1.5-6.5c-.06-.27.11-.54.37-.6c.27-.06.54.11.6.37l1.41 6.11h.17l.93-6.08a.495.495 0 0 1 .98 0l.93 6.08h.17l1.41-6.11a.49.49 0 0 1 .6-.37c.27.06.44.33.37.6l-1.5 6.5c-.05.23-.25.39-.49.39Z"/><path fill="currentColor" d="M12.5 9h-9c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h9c.28 0 .5.22.5.5s-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}