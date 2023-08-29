package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8v32M7 18.2h13.2v11.7H7z"/><circle cx="10.7" cy="21.2" r="1.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.8 29.6l3.3-5l1.9 2l3.1-5.4l3.2 8.5m8.6-11.5h13.2v11.7H27.9z"/><circle cx="31.6" cy="21.2" r="1.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.7 29.6l3.3-5l1.9 2l3.1-5.4l3.2 8.5"/>`),
		g.Group(children),
	)
}