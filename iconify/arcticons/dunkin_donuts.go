package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DunkinDonuts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.987 37.5V25.903L21.67 37.5V25.903M26.444 10.5v7.756a3.842 3.842 0 0 0 7.683 0V10.5m-20.14 11.597V10.5h2.61a5.074 5.074 0 0 1 5.073 5.074v1.45a5.074 5.074 0 0 1-5.074 5.073h-2.61Zm13.182 3.806V37.5m0-4.039l6.233-7.519m0 11.558l-4.775-5.799m7.038-5.929l-.511 2.463"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}