package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UniversalTuner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.202 4.5v23.156A3.202 3.202 0 0 1 24 30.858h0a3.202 3.202 0 0 1-3.202-3.202V4.5M24 30.858v9.296"/><circle cx="24" cy="41.827" r="1.673" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.117 11.114a9.93 9.93 0 0 0-.593 13.694m14.952 0a9.93 9.93 0 0 0-.593-13.693M12.51 6.721a16.293 16.293 0 0 0-.593 22.481m24.165 0a16.292 16.292 0 0 0-.562-22.45"/>`),
		g.Group(children),
	)
}