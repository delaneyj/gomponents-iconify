package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outlook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M267 93h145q4 0 9.5 5t5.5 12l-127 85h-4l-29-18V93zm0 115l27 18q2 1 4 1h3l1-1q-2 1 29-19.5t64-41.5l32-21v153q0 12-6.5 18t-16.5 6H267V208zm-139-39q13 0 20.5 12.5T156 216t-7.5 34t-21.5 12q-13 0-21-12.5T98 216t8-34t22-13zM0 51L251 3v426L0 377V51zm168 218q16-21 16-54t-15.5-53.5T128 141q-26 0-42 21t-16 56q0 32 16 52t41 20t41-21z"/>`),
		g.Group(children),
	)
}