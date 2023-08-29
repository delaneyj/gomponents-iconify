package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#67b0dd" d="M31.908 5C14.286 5 0 14.556 0 26.342c0 6.61 4.495 12.518 11.547 16.432c-.87 2.877-3.372 6.49-9.717 9.89c0 0-1.83 1.235 0 2.746c0 0 12.267.725 20.98-8.609c2.883.573 5.933.886 9.1.886c17.619 0 31.904-9.556 31.904-21.342C63.814 14.556 49.529 5 31.91 5"/><g fill="#fff"><ellipse cx="14.885" cy="27.1" rx="4.572" ry="3.782"/><path d="M31.908 23.316c-2.527 0-4.574 1.692-4.574 3.781c0 2.088 2.047 3.782 4.574 3.782c2.523 0 4.57-1.694 4.57-3.782c0-2.089-2.046-3.781-4.57-3.781"/><ellipse cx="48.643" cy="27.1" rx="4.571" ry="3.782"/></g>`),
		g.Group(children),
	)
}