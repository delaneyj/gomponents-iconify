package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFlipCamera0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M12 11h5l2-4h10l2 4h5v16H12V11Z"/><circle cx="24" cy="18" r="4" fill="#000" stroke="#000"/><path stroke="#fff" d="M24 38C12.954 38 4 33.523 4 28c0-1.422.594-2.775 1.664-4M24 38l-4-4m4 4l-4 4m12-4.832C39.064 35.625 44 32.1 44 28c0-1.422-.594-2.775-1.664-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFlipCamera0)"/>`),
		g.Group(children),
	)
}