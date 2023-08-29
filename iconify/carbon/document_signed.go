package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSigned(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.414 30H15v-4.414l5.034-5.034A4.607 4.607 0 0 1 20 20a5 5 0 1 1 4.448 4.966ZM17 28h1.586l5.206-5.206l.54.124a3.035 3.035 0 1 0-2.25-2.25l.124.54L17 26.414Z"/><circle cx="25" cy="20" r="1" fill="currentColor"/><path fill="currentColor" d="M8 6h12v2H8zm0 4h12v2H8zm0 4h6v2H8zm0 10h4v2H8z"/><path fill="currentColor" d="M12 30H6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v9h-2V4H6v24h6Z"/>`),
		g.Group(children),
	)
}