package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsTwoSd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a1.75 1.75 0 0 0-1.2.51L8.92 18.88a1.76 1.76 0 0 0-.51 1.25h0v21.4a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V20.39a1.82 1.82 0 0 0 0-.57V6.45a2 2 0 0 0-2-1.95H24Zm-8.33 13.94a1.68 1.68 0 0 1 1.17.48l1.38 1.38a9.91 9.91 0 0 1 11.56 0l1.38-1.38a1.64 1.64 0 0 1 2.3 0h0a1.61 1.61 0 0 1 0 2.29l-1.39 1.37a10.06 10.06 0 0 1 1.88 5.8v11a1 1 0 0 1-1 1H15.11a1 1 0 0 1-1-1v-11A9.91 9.91 0 0 1 16 22.57l-1.44-1.37a1.61 1.61 0 0 1 0-2.29h0a1.55 1.55 0 0 1 1.13-.46Zm4.15 7.32a1.83 1.83 0 1 0 1.83 1.83h0a1.83 1.83 0 0 0-1.83-1.83Zm8.36 0a1.83 1.83 0 1 0 0 3.66h0a1.83 1.83 0 1 0 0-3.66Zm-5.61-16v4.23m4.52-7.12v7.12m4.53-7.12v7.12m4.53-7.12v7.12"/>`),
		g.Group(children),
	)
}