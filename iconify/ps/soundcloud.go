package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 238q0 24 18 39v-78Q2 215 2 238zm37-49v98q5 2 15 2h5V187h-5q-5 0-15 2zm43 6q-2 0-4-2v96h19V161q-12 15-15 34zm34-52v146h20V133q-9 3-20 10zm39-14v160h19V130q-4-1-12-1h-7zm49 11q-11-5-20-8v157h29V129q-3 3-9 11zm19-21v170h183q56-3 56-56q0-23-15.5-39.5T408 177q-10 0-22 5q-4-37-32-62t-66-25t-65 24z"/>`),
		g.Group(children),
	)
}