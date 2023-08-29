package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPalm0"><g fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18.333 35.813L8.62 24.933a3.8 3.8 0 1 1 5.803-4.905L16 22V8.994a3 3 0 1 1 6 0V7a3 3 0 1 1 6 0v1.83a3 3 0 1 1 6 0v3.316a3 3 0 0 1 6 0v13.682c0 2.14-.678 4.227-1.937 5.958l-2.833 3.898a.768.768 0 0 1-.621.316H18.75a.56.56 0 0 1-.417-.187Z"/><rect width="16" height="8" x="19" y="36" rx="1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPalm0)"/>`),
		g.Group(children),
	)
}