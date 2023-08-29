package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCircus0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M9 26s-.076.787-1 5c-.98 4.465-4 13-4 13h40s-3.02-8.535-4-13c-.924-4.213-1-5-1-5"/><path stroke-linejoin="round" d="M20 28c1 9-4 16-4 16m12-16c-1 9 4 16 4 16"/><path fill="#fff" stroke-linejoin="round" d="M6 18s7.592-.055 11-1c2.866-.795 7-3 7-3s3.823 2.194 6.5 3c3.572 1.075 11.5 1 11.5 1l-1 7s-4 2-5 2s-3-2-4-2s-3.5 3-4 3s-3-3-4-3s-3 3.5-4 3.5s-3-3.5-4-3.5c-.198 0-.474.078-.79.204c-2.142.852-4.489 1.284-6.629.428L7 25l-1-7Z"/><path d="M24 5v10"/><path stroke-linejoin="round" d="M36 11V5s-1.5 3-6 0s-6 0-6 0v6s1.5-3 6 0s6 0 6 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCircus0)"/>`),
		g.Group(children),
	)
}