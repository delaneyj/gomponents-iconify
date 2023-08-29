package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kerning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1001.523 639h-40q-15 0-31.5-10.5t-24.5-25.5l-30-92h-279l-30 92q-8 15-24.5 25.5t-31.5 10.5h-40q-15 0-20-10.5t3-25.5l186-567q8-15 24.5-25.5t31.5-10.5h81q15 0 31.5 10.5t24.5 25.5l186 567q8 15 3 25.5t-20 10.5zm-266-552l-97 296h195zm-351 516q-8 15-24.5 25.5t-31.5 10.5h-81q-15 0-31.5-10.5t-24.5-25.5l-186-567q-8-15-3-25.5t20-10.5h40q15 0 31.5 10.5t24.5 25.5l170 516l169-516q8-15 24.5-25.5t31.5-10.5h41q15 0 19.5 10.5t-3.5 25.5zm-294 109q8-8 19-8t19 8v56h768v-56q8-8 19-8t19 8l82 66q8 9 8 21.5t-8 21.5l-82 65q-8 9-19 9t-19-9v-54h-768v54q-8 9-19 9t-19-9l-82-65q-8-9-8-21.5t8-21.5z"/>`),
		g.Group(children),
	)
}