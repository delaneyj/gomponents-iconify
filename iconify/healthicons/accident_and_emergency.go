package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccidentAndEmergency(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M20.862 19.532a1 1 0 0 1 .868 1.116l-.904 7.238a1 1 0 0 1-1.985-.248l.905-7.238a1 1 0 0 1 1.116-.868Z"/><path fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm8.454 13.662c.195-1.656 1.392-3.048 3.035-3.332c2.55-.44 4.59-.437 7.041-.008c1.635.286 2.82 1.675 3.015 3.323L32 32H16l1.454-12.338ZM15 36v-2h18v2H15Zm-1.955-11a11.17 11.17 0 0 1 0-2H10v2h3.045Zm.966-5.613a10.98 10.98 0 0 1 1.001-1.73l-2.636-1.523l-1 1.732l2.635 1.521Zm3.645-4.375c.543-.383 1.122-.72 1.732-1.001l-1.522-2.635l-1.732 1l1.522 2.636ZM23 13.045a11.17 11.17 0 0 1 2 0V10h-2v3.045Zm5.613.966a11.03 11.03 0 0 1 1.73 1.001l1.523-2.636l-1.732-1l-1.521 2.635Zm4.375 3.645c.383.543.72 1.122 1.001 1.732l2.635-1.522l-1-1.732l-2.636 1.522ZM34.955 23a11.17 11.17 0 0 1 0 2H38v-2h-3.045Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}