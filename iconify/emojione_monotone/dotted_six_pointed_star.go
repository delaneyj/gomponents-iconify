package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DottedSixPointedStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M50.262 32L60 16H40.522L32 2l-8.521 14H4l9.739 16L4 48h19.479L32 62l8.522-14H60l-9.738-16zm.459-11l-3.547 5.928L43.566 21h7.155zm-6.58 11l-6.578 11H26.439l-6.578-11l6.578-11h11.123l6.579 11zM32 11.701L34.572 16H29.43L32 11.701zM13.283 21h7.151l-3.607 5.926L13.283 21zm0 22l3.544-5.926L20.435 43h-7.152zM32 52.299L29.43 48h5.143L32 52.299zm15.174-15.227L50.721 43h-7.154l3.607-5.928z"/><ellipse cx="32" cy="32" fill="currentColor" rx="4.199" ry="4.286"/>`),
		g.Group(children),
	)
}