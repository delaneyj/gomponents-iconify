package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPillCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="m18.85 13.192l-6.365-6.364a4 4 0 0 0-5.657 5.657l6.364 6.364a4 4 0 1 0 5.657-5.657ZM7.535 7.536a3 3 0 0 1 4.242 0l6.364 6.364a3 3 0 1 1-4.242 4.242l-6.364-6.364a3 3 0 0 1 0-4.242Z" clip-rule="evenodd"/><path d="m16.037 10.58l-.243.97c-1.201-.3-2.223-.154-3.101.432c-.87.58-1.454 1.687-1.73 3.355l-.987-.164c.318-1.917 1.032-3.27 2.162-4.023c1.122-.748 2.434-.936 3.899-.57Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPillCircleFilled0)"/></g>`),
		g.Group(children),
	)
}