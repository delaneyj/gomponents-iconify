package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MongodbOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.369-.338a.5.5 0 0 0-.738 0L7.5.5Zm0 13l-.393.309a.5.5 0 0 0 .786 0L7.5 13.5ZM4.623 9.838l-.393.31l.393-.31Zm.246-6.467L4.5 3.032l.369.337Zm5.262 0l.369-.338l-.369.337Zm.246 6.467l.393.31l-.393-.31ZM8 15V.5H7V15h1Zm-.107-1.809L5.016 9.53l-.786.618l2.877 3.662l.786-.618ZM5.237 3.708L7.87.838L7.13.162L4.5 3.032l.736.676ZM7.131.838l2.632 2.87l.737-.675L7.869.163L7.13.837Zm2.853 8.691l-2.877 3.662l.786.618l2.877-3.662l-.786-.618Zm-.221-5.82a4.5 4.5 0 0 1 .22 5.82l.787.618a5.5 5.5 0 0 0-.27-7.114l-.737.675Zm-4.747 5.82a4.5 4.5 0 0 1 .221-5.82L4.5 3.032a5.5 5.5 0 0 0-.27 7.114l.786-.618Z"/>`),
		g.Group(children),
	)
}