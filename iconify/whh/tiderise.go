package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tiderise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 590q-53-18-118-15t-125.5 20.5t-130 39t-132 40.5T258 700.5T128 695q-54-11-91-35.5T0 608q0-29 38-43t90-4q43 19 103 16t121-20.5t133-39t139-41t139-26t133 5.5q54 11 91 35.5t37 51.5q0 30-39.5 46t-88.5 1zM576 192v95q0 14-9.5 23t-22.5 9h-64q-13 0-22.5-9t-9.5-23v-95H335q-25-23-8-38L489 8q9-8 23-8t24 8l162 146q16 15-9 38H576zM128 881q43 19 103 16t121-20.5t133-39t139-41t139-26t133 5.5q54 11 91 35.5t37 51.5q0 30-39.5 46t-88.5 1q-53-18-118-15t-125.5 20.5t-130 39t-132 40.5t-132.5 25.5t-130-5.5q-54-11-91-35.5T0 927t38-42t90-4z"/>`),
		g.Group(children),
	)
}