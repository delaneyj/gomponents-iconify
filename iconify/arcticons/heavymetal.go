package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heavymetal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 19.173l3.364 3.364h11.772v2.523l-5.466 3.364L9.704 25.9l-.84-3.364M42.5 19.173l-3.364 3.364H27.364v2.523l5.34 3.364l5.592-2.523l.84-3.364m-6.306 5.886v10.932M15.17 28.423v10.932m4.626 2.523l.84-2.523h6.728l.84 2.523"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.704 8.241l9.25 5.046l1.682 9.25M38.295 8.241l-9.25 5.046l-1.681 9.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.954 13.287L24 15.81l5.045-2.523M6.761 10.764l9.25 4.625m4.625 9.671v10.931h6.728V25.06M19.796 6.98L24 9.923l4.204-2.943a10.733 10.733 0 0 0-8.409 0ZM7.182 14.548l9.25 4.625m24.807-8.409l-9.25 4.625m8.829-.841l-9.25 4.625"/>`),
		g.Group(children),
	)
}