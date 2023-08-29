package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quizlet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.394 2.546A21.147 21.147 0 0 0 7.98 9.864a21.35 21.35 0 0 0 7.005 33.24a21.11 21.11 0 0 0 18.872-.406a.588.588 0 0 1 .3-.064a.582.582 0 0 1 .291.097a16.376 16.376 0 0 0 9.312 2.768a.593.593 0 0 0 .593-.595V37.67a.598.598 0 0 0-.137-.389a.586.586 0 0 0-.355-.206a8.395 8.395 0 0 1-1.982-.571a.59.59 0 0 1-.327-.375a.617.617 0 0 1-.018-.255a.603.603 0 0 1 .09-.24a21.366 21.366 0 0 0-5.857-29.56A21.124 21.124 0 0 0 25.37 2.533ZM11.177 23.819a12.909 12.909 0 0 1 7.92-11.937a12.797 12.797 0 0 1 14 2.792a12.944 12.944 0 0 1-1.949 19.865a12.804 12.804 0 0 1-16.213-1.604a12.942 12.942 0 0 1-3.767-9.118Z"/>`),
		g.Group(children),
	)
}