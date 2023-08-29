package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="22" cy="24" r="2" fill="currentColor"/><path fill="none" d="M22 28a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.003 2.003 0 0 0-2-2Z"/><path fill="currentColor" d="M29.777 23.479A8.64 8.64 0 0 0 22 18a8.64 8.64 0 0 0-7.777 5.479L14 24l.223.521A8.64 8.64 0 0 0 22 30a8.64 8.64 0 0 0 7.777-5.479L30 24ZM22 28a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><path fill="currentColor" d="M12 28H7V7h3v3h12V7h3v9h2V7a2 2 0 0 0-2-2h-3V4a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v1H7a2 2 0 0 0-2 2v21a2 2 0 0 0 2 2h5Zm0-24h8v4h-8Z"/>`),
		g.Group(children),
	)
}