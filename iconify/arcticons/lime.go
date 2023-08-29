package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.329 40.011c-4.21 2.43-7.663 2.056-7.703-.336l-.149-9.074a2.577 2.577 0 0 1 3.627-2.095l7.783 4.667c2.06 1.235.652 4.407-3.558 6.838Zm9.201-16.084c0 4.861-2.05 7.664-4.142 6.502l-7.934-4.408a2.578 2.578 0 0 1 0-4.187l7.934-4.409c2.1-1.166 4.143 1.64 4.143 6.502ZM33.202 7.916c4.21 2.43 5.612 5.607 3.56 6.838l-7.785 4.667a2.577 2.577 0 0 1-3.626-2.095l.15-9.074c.04-2.401 3.491-2.767 7.701-.336Zm-18.531.073c4.21-2.43 7.663-2.056 7.702.336l.15 9.075a2.577 2.577 0 0 1-3.627 2.094l-7.783-4.667c-2.06-1.235-.652-4.407 3.558-6.838ZM5.47 24.073c0-4.861 2.05-7.664 4.142-6.502l7.934 4.408a2.578 2.578 0 0 1 0 4.187l-7.934 4.409c-2.1 1.166-4.143-1.64-4.143-6.502Zm9.328 16.011c-4.21-2.43-5.612-5.607-3.56-6.838l7.785-4.666a2.577 2.577 0 0 1 3.626 2.094l-.15 9.074c-.039 2.401-3.491 2.767-7.701.336Z"/><circle cx="24" cy="24" r="2.694" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}