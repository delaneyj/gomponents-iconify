package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1025H256q-27 0-45.5-18.5T192 961v-81q-85-31-138.5-105T0 608t53.5-167T192 337V65q0-26 18.5-45T256 1h384v352q0 13 9.5 22.5T672 385h351l1 1v575q0 26-19 45t-45 19zM632 870L481 720q31-53 31-112q0-92-65.5-158T288 384t-158.5 65.5T64 608t65.5 158.5T288 832q58 0 112-31l150 151q8 8 20 8t21-8l41-41q8-8 8-20t-8-21zM288 768q-66 0-113-46.5t-47-113T175 495t113-47t113 47t47 113.5t-47 113T288 768zM704 0q26 1 44 19l257 257q19 19 18 45H704V0z"/>`),
		g.Group(children),
	)
}