package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1014.188 764l-250 250q-11 10-25.5 10t-25-10.5t-10.5-25t11-24.5l250-250q10-11 24.5-11t25 10.5t10.5 25t-10 25.5zm-322 155q-9 9-22.5 9t-22.5-9l-317-317q-10-10-10-23t10-23l226-226q10-10 23-10t22 10l318 317q9 9 9 22.5t-9 22.5zm-91-493q-9-10-22-10t-23 10l-25 25l244 244q9 9 22.5 9t22.5-9l25-26zm-272 74q-12 12-28.5 12t-28.5-12l-86-85q-29-30-55.5-67.5t-45.5-74.5t-35.5-76t-26-71t-16.5-60.5t-7-43.5t2-20q5-1 20-1.5t43.5 7t60.5 16.5t71 25.5t76 35.5t74.5 45.5t66.5 56.5l86 85q12 12 12 28.5t-12 28.5zm22-250q-75-74-197-127q44 79 96 132l86 85q12 12 28.5 12t28.5-12l24-23z"/>`),
		g.Group(children),
	)
}