package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerprintPhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11V6.362c0-.51.1-1 .284-1.454M22 11V7.815m-7.778-5.08A5.507 5.507 0 0 1 17 2c2.28 0 4.203 1.33 4.805 3.15M15 12V9.824M19 12V6.853C19 5.83 18.105 5 17 5s-2 .83-2 1.853v.794M8 17.01l.01-.011M8 5H3.6a.6.6 0 0 0-.6.6v14.8a.6.6 0 0 0 .6.6h8.8a.6.6 0 0 0 .6-.6V16"/>`),
		g.Group(children),
	)
}