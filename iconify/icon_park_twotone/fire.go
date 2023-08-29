package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFire0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 44c8.235 0 15-6.526 15-14.902c0-2.056-.105-4.26-1.245-7.686c-1.14-3.426-1.369-3.868-2.574-5.984c-.515 4.317-3.27 6.117-3.97 6.655c0-.56-1.666-6.747-4.193-10.45C24.537 8 21.163 5.617 19.185 4c0 3.07-.863 7.634-2.1 9.96c-1.236 2.325-1.468 2.41-3.013 4.14c-1.544 1.73-2.253 2.265-3.545 4.365C9.236 24.565 9 27.362 9 29.418C9 37.794 15.765 44 24 44Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFire0)"/>`),
		g.Group(children),
	)
}